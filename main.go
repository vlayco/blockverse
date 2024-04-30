package main

import (
	"context"
	"log"
	"time"

	"github.com/vlayco/blockverse/node"
	"github.com/vlayco/blockverse/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	makeNode(":3000", []string{})
	time.Sleep(time.Second * 1)
	makeNode(":4000", []string{":3000"})
	time.Sleep(time.Second * 2)
	makeNode(":5000", []string{":4000"})

	// go func() {
	// 	for {
	// 		makeTransaction()
	// 	}
	// }()

	// log.Fatal(node.Start(":3000"))
	select {}
}

func makeNode(listenAddr string, bootstrapNodes []string) *node.Node {
	n := node.NewNode()
	go n.Start(listenAddr, bootstrapNodes)

	return n
}

// just for testing
func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version:    "blockverse-1",
		Height:     1,
		ListenAddr: ":4000",
	}

	_, err = c.Handshake(context.Background(), version)
	if err != nil {
		log.Fatal(err)
	}
}
