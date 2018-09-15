# selpg

a Linux command line utility

Linux命令行使用程序

### 概述

selpg允许用户从输入文本中抽取一定范围的页，这些输入文本可以来自标准输入，文件或另一个进程。页的范围由起始页和终止页决定。在管道中，输入输出和错误流重定向的情况下也可使用该工具。

### 使用selpg

以下说明selpg的使用方法

    $ ./selpg --help

    Usage: ./selpg -s=STARTPAGE -e=ENDPAGE [OPTION]... [FILE]...
    Select specified pages from file or standard input.
    
    With no FILE, or when FILE is -, read standard input.
    
    	-s=STARTPAGE	Pages number starts at STARTPAGE
    	-e=ENDPAGE	Pages number ends at ENDPAGE
    	-l=PAGELENGTH	The number of lines of each page
    	-f	Input file use 'f' to seperate two pages
    	-d	The destination of output

selpg命令的必要参数包括3个：

- 命令名本身
- -s=STARTPAGE
- -e=ENDPAGE

其中，STARTPAGE指定了起始页，ENDPAGE指定了终止页，即抽取页面的范围是[STARTPAGE, ENDPAGE] 。

selpg命令的可选参数包括：

- -l=PAGELENGTH和-f
- -d=DESTINATION

其中，-l和-f是互斥的。selpg可以处理两种输入文本：

1. 该类文本的页行数固定。这是缺省类型，因此不必给出选项进行说明。也就是说，如果既没有给出-l=PAGENUMBER也没有给出-f选项，则 selpg 会理解为页有固定的长度。缺省值为每页 72 行，该缺省值可以用-l=PAGELENGTH选项覆盖。
2. 该类型文本的页由 ASCII 换页字符（十进制数值为 12，在 C 中用’f'表示）定界。该格式与“每页行数固定”格式相比的好处在于，当每页的行数有很大不同而且文件有很多页时，该格式可以节省磁盘空间。在含有文本的行后面，该类型的页只需要一个字符'f'就可以表示该页的结束。

selpg 还允许用户使用-d=DESTINATION来指定程序输出的去处。缺省值为标准输出。

### 设计说明

代码使用GO实现。源码主要包含3个函数：

- setUsage(myArgs *Args)
- argsProcess(myArgs *Args)
- fileProcess(myArgs *Args)

Args是一个结构体，用于储存命令行的所有参数：

```go
type Args struct {
	programName string
	startPage   int
	endPage     int
	srcFile     string
	pageLength  int
	pageType    bool  // true for using '\f', false for -l=Number
	desProgram  string
}
```

**setUsage(myArgs \*Args):**

该函数主要使用了flag包来实现。它的功能主要有两个：

1. 为`flag.Usage`赋值。`Usage`的主要作用是提供命令的使用说明，即默认情况下输入

   ```
   $ ./selpg --help
   ```

   时，得到的输出。

2. 完成对所有参数的解析。`flag.XXXVar`提供了十分方便的功能，可以自动解析用户输入命令中的参数（以`-`开头的），并完成赋值，例如：

   ```go
   flag.IntVar(&myArgs.startPage, "s", -1, "specify start page.")
   ```

   这里，若输入命令中含有`-s=3`，函数会自动以Int类型将”3“这个值，赋给`myArgs.startPage`，且规定了，当不含`-s`参数时，缺省值为-1。

**argsProcess(myArgs \*Args):**

该函数主要负责对参数进行逻辑判断和处理。包括：

1. `-s`和`-e`为强制选项，用户输入中必须包含这两个参数，否则报错。
2. 起始页的数值要小于或等于终止页，否则报错。
3. 若用户已经使用了`-f`参数，则不能再使用`-l`参数，否则报错。
4. 若用户没有使用`-f`和`-l`参数，则页面长度设置为默认长度。该程序中为72。

**fileProcess(myArgs \*Args)**

负责根据命令和参数，进行文本的读写。大部分的功能可以通过直接使用包来实现：

- os：用于打开文件
- bufio：用于文件流读写
- io：用于检测文件结束符
- fmt：用于格式化输出

为了提高代码可读性，实现了两个函数方便调用：

    readByPage(inputReader *bufio.Reader, myArgs *Args)
    readByLine(inputReader *bufio.Reader, myArgs *Args)

分别用于针对上述的两种不同类型的文本进行读取。

`fileProcess`函数伪代码如下：

```clike
获取srcFile值
if srcFile 为空 then
	输入为标准输入
	if 页面以'\f'分隔 then
		readByPage()
	else then
		readByLine()
else than
	输入来自于文件srcFile
	打开srcFile，获取输入流
	if 页面以'\f'分隔 then
		readByPage()
	else then
		readByLine()
```

**readByPage(inputReader \*bufio.Reader, myArgs \*Args)**

该函数用于读取以'\f'作为分隔符的文本。

在函数中，会遍历读取整个文件的每一页，以`pageCount`记录当前页数，当`pageCount`满足用户的抽取范围时，将当前页内容输出到屏幕或目标程序/文件。当读取到`EOF`，则会输出剩余内容并跳出循环。

为了验证`-d`的功能，这里默认为会将`./out`程序作为`-d`的参数值。`./out`的作用是在每个输入的字符串后加上"success!"。

```go
// 打开./go的输入管道，将该程序输出传输到管道
cmd := exec.Command("./out")          // 创建命令"./out"
echoInPipe, err := cmd.StdinPipe()    // 打开./out的标准输入管道
check(err)                            // 错误检测
echoInPipe.Write([]byte(page + "\n")) // 向管道中写入文本
echoInPipe.Close()                    // 关闭管道
cmd.Stdout = os.Stdout                // ./out将会输出到屏幕
cmd.Run()                             // 运行./out命令
```

同时，该函数会判断起始页码和终止页码与总页数的关系，当不符合逻辑时会给出相应的警告。

**readByLine(inputReader \*bufio.Reader, myArgs \*Args)**

功能与实现基本和`readByPage`相同，只是处理的文本类型不同。这里以规定的行数来分隔两页。计数的方法则是采用`lineCount`，来判断当前页数。



另外，为了处理大量的相同代码，在错误检测和错误输出处理上，实现了两个函数：

- check(err error)
- processError(name string, errorStr string)



### 使用与测试

    $ ./selpg -s=2 -e=5

该命令只使用了所有强制选项，缺省情况下，从标准输入读取（输入以ctrl+c结束），抽取其中2到5页，并缺省地输出到标准输出（命令行）。

    $ ./selpg -s=2 -e=5 in.txt

该命令由in.txt中读取2到5页，并输出到命令行。

    $ ./selpg -s=2 -s=5 < in.txt

该命令读取标准输入，但标准输入已被重定向为来自in.txt，从中读取2到5页，并输出到命令行。

    $ ./in | ./selpg -s=2 -e=3

这里的./in，用于输出1000行文本，每一行内容为"lineX"，其中X为行号，例如第30行内容为”line30“。这一命令中，./in的标准输出被重定向至./selpg的标准输入。将第2页到第3页的内容输出到屏幕。测试中，输出内容为”line72”到“line216”，正确。

    $ ./selpg -s=2 -e=3 in.txt >out.txt

selpg将in.txt的第2页到第3页写到标准输出，标准输出被重定向至out.txt。out.txt中的内容为“line72”到“line216”。结果正确。

    $ ./selpg -s=2 -e=3 in.txt 2>error.txt
    $ ./selpg -s=10 -e=3 in.txt 2>error.txt

selpg 将第 2 页到第 3 页写至标准输出（屏幕）。所有的错误消息被 shell／内核重定向至error.txt。由于第一行命令中没有错误，error.txt为空。使用第二行命令后，可以看到error.txt的内容为：“./selpg: STARTPAGE(10) should not be greater than ENDPAGE(3).”

    $ ./selpg -s=10 -e=3 in.txt >out.txt 2>error.txt

所有的错误消息被 shell／内核重定向至error.txt。可以看到error.txt的内容为：“./selpg: STARTPAGE(10) should not be greater than ENDPAGE(3).”。而out.txt内容为Usage的内容，因为在发生这样的错误时程序会终止并输出Usage。

    $ ./selpg -s=10 -e=3 in.txt >out.txt 2>error.txt

所有输出被废弃，而错误信息输出在屏幕上：”./selpg: STARTPAGE(10) should not be greater than ENDPAGE(3).“

    $ ./selpg -s=2 -e=5 in.txt | ./out

 从in.txt读取第2页到第5页，输出重定向到程序./out中，成为它的标准输入。这里，./out的功能是在每个输入的string后面加上"success!"。可以看到，屏幕输出为”line73success!“到”line360success!“，结果正确。

    $ ./selpg -s=2 -e=5 -l=50 in.txt

从in.txt读取第2页到第5页，每页长度设定为50行，输出到屏幕。可以看到，屏幕输出为“line51"到”line250“，结果正确。

    $ ./selpg -s=2 -e=5 -f in.txt

从in.txt读取第2页到第5页，以'\f'为换页符，输出到屏幕。

    $ ./selpg -s=2 -e=5 -d=./out in.txt 

由于没有实际打印机，测试时，将页通过管道输入到程序`./out`的标准输入（程序功能见上文），这里第2至5页输送到`./out`命令。屏幕上输出为”line73success!“到”line360success!“，结果正确。
