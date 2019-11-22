package consensus

import "github.com/MatrixAINetwork/go-matrix/sharding/types"

type ConsensusInterface interface {
	Seal(blockInterface types.BlockInterface) (types.BlockInterface,error)
}
