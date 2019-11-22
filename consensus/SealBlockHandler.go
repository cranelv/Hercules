package consensus

import (
	"github.com/MatrixAINetwork/go-matrix/sharding/types"
)

type SealBlockHandler struct {
	blockChan chan types.BlockInterface
	consensus ConsensusInterface
}
func (seal *SealBlockHandler)Channel() interface{} {
	return seal.blockChan
}
func (seal *SealBlockHandler)Handler()types.HandleRequestInterface{
	return seal
}

func(seal *SealBlockHandler)Handle(Data interface{}){
	if Data != nil{
		block := Data.(types.BlockInterface)
		if block != nil{
			err := seal.consensus.Seal(block)
			if err != nil {

			}
		}
	}
}

