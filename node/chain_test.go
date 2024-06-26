package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vlayco/blockverse/crypto"
	"github.com/vlayco/blockverse/proto"
	"github.com/vlayco/blockverse/types"
	"github.com/vlayco/blockverse/util"
)

func randomBlock(t *testing.T, chain *Chain) *proto.Block {
	privKey := crypto.GeneratePrivateKey()
	b := util.RandomBlock()
	prevBlock, err := chain.GetBlockByHeight(chain.Height())
	require.NoError(t, err)
	b.Header.PrevHash = types.HashBlock(prevBlock)
	types.SignBlock(privKey, b)
	return b
}

func TestNewChain(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())
	assert.Equal(t, 0, chain.Height())
	_, err := chain.GetBlockByHeight(0)
	assert.NoError(t, err)
}

func TestChainHeight(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())

	for i := 0; i < 1000; i++ {
		b := randomBlock(t, chain)

		// require stops test execution with first error occurrence.
		require.NoError(t, chain.AddBlock(b))
		require.Equal(t, chain.Height(), i+1) // because of added genesisi block.
	}
}

func TestAddBlock(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())

	for i := 0; i < 100; i++ {
		block := randomBlock(t, chain)

		blockHash := types.HashBlock(block)

		require.NoError(t, chain.AddBlock(block))

		fetchedBlock, err := chain.GetBlockByHash(blockHash)
		require.Nil(t, err)
		require.Equal(t, block, fetchedBlock)

		fetchedBlockByHeight, err := chain.GetBlockByHeight(i + 1)
		require.NoError(t, err)
		require.Equal(t, block, fetchedBlockByHeight)
	}

}
