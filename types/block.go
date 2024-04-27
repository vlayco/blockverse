package types

import (
	"crypto/sha256"

	"github.com/vlayco/blockverse/crypto"
	"github.com/vlayco/blockverse/proto"
	pb "google.golang.org/protobuf/proto"
)

func SignBlock(pk *crypto.PrivateKey, b *proto.Block) *crypto.Signature {
	return pk.Sign(HashBlock(b))
}

// HashBlock returns sha256 of the header.
func HashBlock(block *proto.Block) []byte {
	bytes, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(bytes)
	return hash[:]
}
