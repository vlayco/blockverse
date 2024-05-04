package main

import (
	"context"
	"log"
	"time"

	"github.com/vlayco/blockverse/crypto"
	"github.com/vlayco/blockverse/node"
	"github.com/vlayco/blockverse/proto"
	"github.com/vlayco/blockverse/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	makeNode(":3000", []string{}, true)
	time.Sleep(time.Second * 1)
	makeNode(":4000", []string{":3000"}, false)
	time.Sleep(time.Second * 2)
	makeNode(":5000", []string{":4000"}, false)

	// go func() {
	// 	for {
	// 		makeTransaction()
	// 	}
	// }()

	// log.Fatal(node.Start(":3000"))
	for {
		time.Sleep(time.Millisecond * 100)
		makeTransaction()
	}
	// select {}
}

func makeNode(listenAddr string, bootstrapNodes []string, isValidator bool) *node.Node {
	cfg := node.ServerConfig{
		Version:    "blockverse-1",
		ListenAddr: listenAddr,
	}

	if isValidator {
		cfg.PrivateKey = crypto.GeneratePrivateKey()
	}

	n := node.NewNode(cfg)
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

	privKey := crypto.GeneratePrivateKey()
	tx := &proto.Transaction{
		Version: 1,
		Inputs: []*proto.TxInput{
			{
				PrevTxHash:   util.RandomHash(),
				PrevOutIndex: 0,
				PublicKey:    privKey.Public().Bytes(),
			},
		},
		Outputs: []*proto.TxOutput{
			{
				Amount:  99,
				Address: privKey.Public().Address().Bytes(),
			},
		},
	}

	_, err = c.HandleTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
}
