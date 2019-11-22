package blockService

import "github.com/cranelv/Hercules/common/types"

type BlockHandleInterface interface {
	IntegrateBlock(types.HeaderInterface,types.TxBlockInterface) types.BlockInterface
	SeparateBlock(blockInterface types.BlockInterface)(types.HeaderInterface,types.TxBlockInterface)
}