package node

import (
	"context"
	"fmt"

	"github.com/vlayco/blockverse/proto"
	"google.golang.org/grpc/peer"
)

type Node struct {
	// When we handle the transaction we want to broadcast it to all the peers on
	// the network.
	// peers map[net.Addr]*grpc.ClientConn

	version string
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		version: "blockverse-1",
	}
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	ourVersion := &proto.Version{
		Version: n.version,
		Height:  100,
	}

	peer, _ := peer.FromContext(ctx)

	fmt.Printf("received version from %s: %+v\n", v, peer.Addr)

	return ourVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("received tx from: ", peer)
	return &proto.Ack{}, nil
}
