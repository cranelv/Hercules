package pipe

import "github.com/MatrixAINetwork/go-matrix/sharding/types"

type HeaderHandlePipe struct {
	pipe PipeInterface
}
func (hp *HeaderHandlePipe)Connect()error{
	return hp.pipe.Connect()
}
func (hp *HeaderHandlePipe)GenerateHeader(parent types.HeaderInterface,txBlock types.TxBlockInterface,callback func(interface{},error)){
	hp.pipe.Request(Message{HEADER_GenerateHeader,[]interface{}{parent,txBlock}},callback)
}
