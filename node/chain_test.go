package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vlayco/blockverse/types"
	"github.com/vlayco/blockverse/util"
)

func TestChainHeight(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())

	for i := 0; i < 1000; i++ {
		b := util.RandomBlock()
		assert.NoError(t, chain.AddBlock(b))
		assert.Equal(t, chain.Height(), i)
	}
}

func TestAddBlock(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())

	for i := 0; i < 100; i++ {
		var (
			block     = util.RandomBlock()
			blockHash = types.HashBlock(block)
		)
		assert.NoError(t, chain.AddBlock(block))

		fetchedBlock, err := chain.GetBlockByHash(blockHash)
		assert.Nil(t, err)
		assert.Equal(t, block, fetchedBlock)

		fetchedBlockByHeight, err := chain.GetBlockByHeight(i)
		assert.NoError(t, err)
		assert.Equal(t, block, fetchedBlockByHeight)
	}

}
