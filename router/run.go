/**
  @author: Zero
  @date: 2023/3/2 21:19:21
  @desc:

**/

package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	. "summer/constant"
	"syscall"
	"time"
)

// Run 启动项目
func Run() {
	// 设置全局gin环境
	gin.SetMode(Config.Server.Env)
	//将本次运行后的gin日志同时写到控制台和文件
	file, err := os.Create("logs/application.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout, file)

	// 创建服务实例
	application := gin.Default()
	// 使用自定义日志中间件
	//application.Use(middleware.Logger())
	// 绑定路由
	InitRouters(application)
	server := Config.Server
	// 启动服务
	RunHttpServer(application, server.Host, server.Port)
}

// RunHttpServer 	优雅启动一个Gin服务
// engine 	  要启动的服务
// host: 	  服务地址
// port: 	  服务进程端口
func RunHttpServer(engine *gin.Engine, host string, port int) {
	// 通过原生httpServer + gin引擎开启服务
	server := &http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: engine,
	}
	// 这里通过一个协程来开启 http服务,保证服务能够主动性的优雅停止
	go func() {
		Log.Infof("Summer Server Running Success Port In %s", server.Addr)
		// 启动
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Log.Panicf("Summer Server Running Fail: %s", err)
		}
	}()
	// 由于服务由一个协程开启的,所以主协程这里我们需要手动阻塞一下,直到接受到停止信号
	// 创建一个接收操作系统信号的通道
	exit := make(chan os.Signal)
	// 这里表示如果接收到了SIGINT或者SIGTERM系统信号,则会把信号向exit通道发送.
	// syscall.SIGINT: 		用户发送INTR字符,例如在终端执行(Ctrl+C) 触发 kill -2 pid然后进程结束
	// syscall.SIGTERM: 	结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞,直到接收到两种信号其中一种...
	<-exit
	// 信号接收到后,需要一定的时间释放相应的资源。 这里延迟3秒,模拟释放资源
	closeCtx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	// 停止服务
	if err := server.Shutdown(closeCtx); err != nil {
		//TODO 释放资源
		Log.Error(err)
	}
	// 等待释放资源,结束程序.
	select {
	case <-closeCtx.Done():
		Log.Info("Wait Close Resource Timeout...")
	}
	Log.Info("HTTP Server Shutdown Success")
}
