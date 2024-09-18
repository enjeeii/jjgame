package service

import (
	"context"
	"jjgame/internal/logger"
	"jjgame/internal/protoc/hello"
)

type HelloService struct {
	hello.UnimplementedTesterServer
}

func (s HelloService) SayHello(ctx context.Context, msg *hello.HelloReq) (*hello.HelloResp, error) {
	logger.INFO_MSG(msg.Msg)
	resp := hello.HelloResp{
		Msg: "Hello " + msg.Msg,
	}
	return &resp, nil
}
