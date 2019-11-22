package consensus

import "github.com/cranelv/Hercules/common/types"

type ConsensusInterface interface {
	Seal(blockInterface types.BlockInterface) (types.BlockInterface,error)
}
