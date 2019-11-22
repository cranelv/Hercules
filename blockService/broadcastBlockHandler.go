package blockService

import (
	"github.com/MatrixAINetwork/go-matrix/sharding/pipe"
	"github.com/MatrixAINetwork/go-matrix/sharding/types"
)

type BroadcastBlockHandler struct {
	blockChan chan types.BlockInterface
	gateway pipe.GatewayPipe
}
func (broadcast* BroadcastBlockHandler)Channel() interface{} {
	return broadcast.blockChan
}
func (broadcast* BroadcastBlockHandler)Handler()types.HandleRequestInterface{
	return broadcast
}

func(broadcast* BroadcastBlockHandler)Handle(Data interface{}){
	if Data != nil {
		block := Data.(types.BlockInterface)
		if block == nil {
			return
		}
		broadcast.gateway.BroadcastMssage(block,nil, func(result interface{}, err error) {

		})
	}
}

