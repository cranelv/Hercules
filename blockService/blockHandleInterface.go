package blockService

import "github.com/MatrixAINetwork/go-matrix/sharding/types"

type BlockHandleInterface interface {
	IntegrateBlock(types.HeaderInterface,types.TxBlockInterface) types.BlockInterface
	SeparateBlock(blockInterface types.BlockInterface)(types.HeaderInterface,types.TxBlockInterface)
}