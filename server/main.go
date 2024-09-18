package main

import (
	"google.golang.org/grpc"
	"jjgame/internal/config"
	"jjgame/internal/logger"
	"jjgame/internal/protoc/hello"
	"jjgame/internal/service"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 配置管理
	config.InitAppConfig()
	// 日志管理
	logger.InitAppLogger(config.GetLoggerConfig())

	// 启动服务
	var opts []grpc.ServerOption
	gameServer := grpc.NewServer(opts...)

	protocol := config.GetServerProtocol()
	address := config.GetServerAddress()

	lis, err := net.Listen(protocol, address)
	if err != nil {
		logger.PANIC_MSG("Failed to listen:", err)
	}

	hello.RegisterTesterServer(gameServer, service.HelloService{})
	logger.INFO_MSG("Start Serve...")
	go gameServer.Serve(lis)
	logger.INFO_MSG_F("Created listener: %s(%s)", address, protocol)

	// 信号处理
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		select {
		case sig := <-sigCh:
			logger.INFO_MSG("Game Singal:", sig.String())
			gameServer.Stop()
			os.Exit(0)
		}
	}
}
