package main

import (
	"os"
	"web/cloudgo/server"

	"github.com/spf13/pflag"
)

const (
	PORT string = "8080"
)

/* 解析参数，设置端口
 * 并把端口号传递到server中
 * 由server发起监听
 */
func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := pflag.StringP("port", "p", PORT, "PORT for httpd listening")
	pflag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server.Run(port)
}
