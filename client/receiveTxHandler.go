package client

import "github.com/cranelv/Hercules/common/types"

type ReceiveTxHandler struct {
	txchan chan interface{}
	client ClientInterface
}
func (ins *ReceiveTxHandler)Channel() interface{} {
	return ins.txchan
}
func (ins *ReceiveTxHandler)Handler()types.HandleRequestInterface{
	return ins
}

func(ins *ReceiveTxHandler)Handle(Data interface{}){
	if Data != nil{
		tx := Data.(types.TransactionInterface)
		if tx == nil{
			return
		}
		ins.client.SendTx(tx)
	}
}