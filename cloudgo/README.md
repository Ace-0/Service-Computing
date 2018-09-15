# cloudgo

A simple http server written with **Go**

### 运行服务器：

```shell
go run main.go -p 9090
```

### curl命令测试：

```shell
curl -v http://localhost:9090/
```

```shell
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9090 (#0)
> GET / HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.52.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 15 Nov 2017 06:42:52 GMT
< Content-Length: 5
< Content-Type: text/plain; charset=utf-8
< 
* Curl_http_done: called premature == 0
* Connection #0 to host localhost left intact
hello
```

```shell
curl -v http://localhost:9090/time
```

```shell
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9090 (#0)
> GET /time HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.52.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 15 Nov 2017 06:46:57 GMT
< Content-Length: 48
< Content-Type: text/plain; charset=utf-8
< 
Now the time is:
* Curl_http_done: called premature == 0
* Connection #0 to host localhost left intact
14:46:57, 2017-11-15, Wednesday
```

```shell
curl -v http://localhost:9090/?username=Jarvis\&query=golang
```

命令中的`&`要使用转义字符的形式`\&`，否则无法被读取。

```shell
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9090 (#0)
> GET /?username=Jarvis&query=golang HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.52.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 15 Nov 2017 06:48:21 GMT
< Content-Length: 48
< Content-Type: text/plain; charset=utf-8
< 
hello, [Jarvis]
Are you searching for [golang]?
* Curl_http_done: called premature == 0
* Connection #0 to host localhost left intact
```

### ApacheBench压力测试：

```shell
ab -n 1000 -c 100 http://localhost:9090/
```

```shell
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            9090

Document Path:          /
Document Length:        5 bytes

Concurrency Level:      100
Time taken for tests:   0.073 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      121000 bytes
HTML transferred:       5000 bytes
Requests per second:    13733.43 [#/sec] (mean)
Time per request:       7.282 [ms] (mean)
Time per request:       0.073 [ms] (mean, across all concurrent requests)
Transfer rate:          1622.80 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   1.4      2       8
Processing:     0    5   7.3      2      29
Waiting:        0    4   6.3      2      29
Total:          2    7   7.4      4      31

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      5
  75%      6
  80%      8
  90%     24
  95%     28
  98%     30
  99%     30
 100%     31 (longest request)
```

```shell
ab -n 1000 -c 100 http://localhost:8080/time
```

执行的请求数量为1000，并发请求个数为100

结果如下：

```shell
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /time
Document Length:        31 bytes

Concurrency Level:      100
Time taken for tests:   0.036 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      148000 bytes
HTML transferred:       31000 bytes
Requests per second:    27687.03 [#/sec] (mean)
Time per request:       3.612 [ms] (mean)
Time per request:       0.036 [ms] (mean, across all concurrent requests)
Transfer rate:          4001.64 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1    1   0.4      1       2
Processing:     0    2   0.7      2       4
Waiting:        0    2   0.6      2       3
Total:          2    3   0.7      3       6

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      4
  75%      4
  80%      4
  90%      4
  95%      5
  98%      5
  99%      5
 100%      6 (longest request)
```

- 吞吐率(Request per second)：27687.03 [requests/sec]
- 用户平均等待时间(Time per request)：3.612 [ms]
- 服务器平均等待时间(Time per request:across all concurrent requests)：0.036 [ms]