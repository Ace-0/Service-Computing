package mux

import (
	"fmt"
	"net/http"
	"time"
)

/* MyMux 自定义路由
 * 通过实现ServeHTTP来实现接口Handler
 * 最终拦截DefaultServeMux，实现自定义路由的目的
 */
type MyMux struct{}

/* 对应路径"/"的Handler */
func helloUser(w http.ResponseWriter, r *http.Request) {
	/* 默认情况下是不会自动解析的
	 * 要调用request.ParseForm（）完成参数解析
	 */
	r.ParseForm()
	username := r.Form["username"]
	query := r.Form["query"]
	/* 输出到客户端 */
	fmt.Fprintf(w, "hello")
	if username != nil {
		fmt.Fprintf(w, ", %s", username)
	}
	if query != nil {
		fmt.Fprintf(w, "\nAre you searching for %s?\n", query)
	}
}

/* 对应路径"/time"的Handler */
func getTime(w http.ResponseWriter, r *http.Request) {
	timestr := time.Now().Format("15:04:05, 2006-01-02, Monday")
	fmt.Fprintf(w, "Now the time is:\n%s", timestr)
}

/* 实现ServeHTTP以实现接口
 * 根据request选择handler
 * 如果没有路由满足，调用NotFoundHandler的ServeHTTP
 */
func (mux *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		helloUser(w, r)
	} else if r.URL.Path == "/time" {
		getTime(w, r)
	} else {
		http.NotFound(w, r)
	}
}
