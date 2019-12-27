package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"github.com/authhwang/gin-admin/internal/app"
	"github.com/authhwang/gin-admin/pkg/logger"
	"github.com/authhwang/gin-admin/pkg/util"
)

// VERSION 版本号，
// 可以通过编译的方式指定版本号：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "5.0.0"

var (
	configFile string
	modelFile  string
	wwwDir     string
	swaggerDir string
)

func init() {
	flag.StringVar(&configFile, "c", "", "配置文件(.json,.yaml,.toml)")
	flag.StringVar(&modelFile, "m", "", "Casbin的访问控制模型(.conf)")
	flag.StringVar(&wwwDir, "www", "", "静态站点目录")
	flag.StringVar(&swaggerDir, "swagger", "", "swagger目录")
}

func main() {
	flag.Parse()

	if configFile == "" {
		panic("请使用-c指定配置文件")
	}

	var state int32 = 1
	sc := make(chan os.Signal)
	//作用：优雅关闭程序
	//根据从外界收到4个信号，收到后再进行处理
	//syscall.SIGHUP: 终端控制进程结束(终端连接断开)
	//syscall.SIGINT: 用户发送INTR字符(Ctrl+C)触发
	//syscall.SIGTERM: 结束程序(可以被捕获、阻塞或忽略)
	//syscall.SIGQUIT: 用户发送QUIT字符(Ctrl+/)触发
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := logger.NewTraceIDContext(context.Background(), util.MustUUID())
	span := logger.StartSpanWithCall(ctx)

	call := app.Init(ctx,
		app.SetConfigFile(configFile),
		app.SetModelFile(modelFile),
		app.SetWWWDir(wwwDir),
		app.SetSwaggerDir(swaggerDir),
		app.SetVersion(VERSION))

	select {
	case sig := <-sc:
		atomic.StoreInt32(&state, 0)
		span().Printf("获取到退出信号[%s]", sig.String())
	}

	if call != nil {
		call()
	}
	span().Printf("服务退出")

	os.Exit(int(atomic.LoadInt32(&state)))
}
