package blockChainService

import "github.com/cranelv/Hercules/common/types"

type BlockChainInterface interface {
	InsertBlock(block types.BlockInterface)error
	InsertHeader(header types.HeaderInterface)error
}
