package runningSpace

import (
	"github.com/MatrixAINetwork/go-matrix/sharding/types"
)

type TxRunningHandler struct {
	TxChan chan interface{}
	handler RunningSpaceHandler
}

func (tr *TxRunningHandler)Channel() interface{} {
	return tr.TxChan
}
func (tr *TxRunningHandler)Handler()types.HandleRequestInterface{
	return tr
}

func(tr *TxRunningHandler)Handle(Data interface{}){
	if Data != nil{
		input := Data.([]interface{})
		if len(input) == 2{
			name := input[0].(string)
			tx := input[1].([]types.TransactionInterface)
			runningSpace := tr.handler.GetRunningSpace(name)
			if runningSpace != nil{
				runningSpace.CallTransactions(tx)
			}
		}
	}
}

