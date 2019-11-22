package txPool

import (
	"github.com/MatrixAINetwork/go-matrix/sharding/pipe"
	"github.com/MatrixAINetwork/go-matrix/sharding/types"
	"github.com/MatrixAINetwork/go-matrix/sharding/protocol"
)

type ReceiveTxHandler struct {
	TxChan chan interface{}
	rsPipe pipe.RunningSpacePipe
	TxPoolHandle TxPoolInterface
}

func (rt *ReceiveTxHandler)Channel() interface{} {
	return rt.TxChan
}
func (rt *ReceiveTxHandler)Handler()types.HandleRequestInterface{
	return rt
}

func(rt *ReceiveTxHandler)Handle(Data interface{}){
	if Data != nil{
		tx := Data.(types.TransactionInterface)
		if tx != nil{
			rt.rsPipe.CreateRunningSpaceByNumber(protocol.BlockNumber_Pending,pipe.RunningSpace_Writable, func(result interface{}, err error) {
				if err != nil{
					// return error
				}
				rsName := result.(string)
				rt.rsPipe.TxCall(rsName,tx, func(result interface{}, err error) {
					if err != nil {

					}else{
						rt.TxPoolHandle.Append(tx)
					}
				})
			})
		}
	}
}
