package hello

import (
	"context"
	"io"
	"jjgame/internal/logger"
	"strconv"
)

type HelloService struct {
	UnimplementedTesterServer
}

// Simple RPC
func (s HelloService) SayHello(ctx context.Context, msg *HelloReq) (*HelloResp, error) {
	logger.INFO_MSG(msg.Msg)
	resp := HelloResp{
		Msg: "Reply " + msg.Msg,
	}
	return &resp, nil
}

// Server-side streaming RPC
func (s HelloService) SayHello2(msg *HelloReq, stream Tester_SayHello2Server) error {
	logger.INFO_MSG(msg.Msg)
	for i := 1; i < 3; i++ {
		resp := HelloResp{
			Msg: "Reply " + msg.Msg + strconv.Itoa(i),
		}
		if err := stream.Send(&resp); err != nil {
			return err
		}
	}
	return nil
}

// Client-side streaming RPC
func (s HelloService) SayHello3(stream Tester_SayHello3Server) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			resp := HelloResp{
				Msg: "Reply client-side rpc.",
			}
			return stream.SendAndClose(&resp)
		} else if err != nil {
			return err
		}
		logger.INFO_MSG(msg.Msg)
	}
}

// Bidirectional streaming RPC
func (s HelloService) SayHello4(stream Tester_SayHello4Server) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		logger.INFO_MSG(msg.Msg)
		for i := 1; i < 3; i++ {
			resp := HelloResp{
				Msg: "Reply " + msg.Msg + strconv.Itoa(i),
			}
			if err := stream.Send(&resp); err != nil {
				return err
			}
		}
	}
}
