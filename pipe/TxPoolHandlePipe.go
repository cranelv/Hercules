package pipe

import "github.com/cranelv/Hercules/common/types"

type TxPoolHandlePipe struct {
	pipe PipeInterface
}
func (tp *TxPoolHandlePipe)Connect()error{
	return tp.pipe.Connect()
}
func (tp *TxPoolHandlePipe)AppendTx(tx types.TransactionInterface,callback func(interface{},error)){
	tp.pipe.Request(Message{TXPOOL_AppendTx,tx},callback)
}
func (tp *TxPoolHandlePipe)GenerateBlock(header types.HeaderInterface,callback func(interface{},error)){
	tp.pipe.Request(Message{TXPOOL_GenerateBlock,header},callback)
}
