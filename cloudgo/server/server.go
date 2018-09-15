package server

import (
	"net/http"
	"web/cloudgo/mux"
)

/* 通过一个ListenAndServe监听服务
 * 底层处理：初始化一个server对象，
 * 然后调用 net.Listen("tcp", addr), 监控设置的端口port。
 * 监控端口之后，调用 srv.Serve(net.Listener) 函数，处理接收客户端的请求信息。
 * 首先通过Listener 接收请求，然后创建一个Conn，
 * 最后单独开了一个goroutine，go c.serve()，把这个请求的数据当做参数扔给这个conn去服务。
 * 用户的每一次请求都是在一个新的goroutine去服务，相互不影响。
 */
func Run(port string) {
	mux := &mux.MyMux{}
	http.ListenAndServe(":"+port, mux)
}
