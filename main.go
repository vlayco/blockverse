package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/vlayco/blockverse/node"
	"github.com/vlayco/blockverse/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	node := node.NewNode()
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	proto.RegisterNodeServer(grpcServer, node)
	fmt.Println("node running on port: ", ":3000")

	go func() {
		for {
			time.Sleep(time.Second * 2)
			makeTransaction()
		}
	}()

	grpcServer.Serve(ln)

}

// just for testing
func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version: "blockverse-1",
		Height:  1,
	}

	_, err = c.Handshake(context.Background(), version)
	if err != nil {
		log.Fatal(err)
	}
}
