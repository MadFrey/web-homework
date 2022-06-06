package main

import (
	login "client/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := login.NewSearchServiceClient(conn)

	for {
		//这段不重要
		fmt.Println("input username&password:")
		iptName := ""
		_, _ = fmt.Scanln(&iptName)
		iptPassword := ""
		_, _ = fmt.Scanln(&iptPassword)

		loginResp, err := c.Search(context.Background(), &login.LoginReq{
			Username: iptName,
			Password: iptPassword,
		})
		if err != nil {
			log.Println(err)
		}

		if loginResp.OK {
			fmt.Println("success")
			break
		}
		fmt.Println("retry")
	}
}
