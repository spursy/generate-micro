package main

import (
	"context"
	"fmt"
	pb "git.5th.im/long-bridge-algo/golang/algo-dispatcher/proto/demo"
	"github.com/google/uuid"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"time"
)

func main() {
	address := "127.0.0.1:8500"
	register := consul.NewRegistry(registry.Addrs(address))
	gCli := grpc.NewClient(client.Registry(register))

	fmt.Println("000000000000")
	cli := pb.NewAlgoDispatcherService("lb.algo.demo", gCli)
	fmt.Println("1111111111")
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"x-request-id": uuid.New().String(),
		"member-id":    "1355",
	})
	fmt.Println("2222222222")
	res, _ := cli.GetDemo(ctx, &pb.DemoRequest{
		K: "1",
	}, client.WithRequestTimeout(5*time.Second), client.WithRetries(2))

	// res, _ := cli.GetUsers(ctx, &servicePb.UserRequest{
	// 	K: "123",
	// })

	fmt.Printf("--- res %v", res)
}
