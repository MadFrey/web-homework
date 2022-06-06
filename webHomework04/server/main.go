package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	login "server/pb"
)

const (
	port = ":50051"
)

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() //获取新服务示例
	login.RegisterSearchServiceServer(s, &server{})

	// 开始处理
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	login.UnimplementedSearchServiceServer
}

func (s *server) Search(ctx context.Context, req *login.LoginReq) (*login.LoginRes, error) {
	resp := &login.LoginRes{}
	log.Println("recv:", req.Username, req.Password)
	if req.Password != GetPassWord(req.Username) {
		resp.OK = false
		return resp, nil
	}
	resp.OK = true
	return resp, nil
}

func GetPassWord(userName string) (password string) {
	return userName + "123456"
}
