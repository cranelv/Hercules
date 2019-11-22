package blockChainService

import "github.com/cranelv/Hercules/common/types"

type InsertHeaderHandler struct {
	headChan chan types.HeaderInterface
	blockchain BlockChainInterface
}
func (ins *InsertHeaderHandler)Channel() interface{} {
	return ins.headChan
}
func (ins *InsertHeaderHandler)Handler()types.HandleRequestInterface{
	return ins
}

func(ins *InsertHeaderHandler)Handle(Data interface{}){
	if Data != nil{
		header := Data.(types.HeaderInterface)
		if header == nil{
			return
		}
		ins.blockchain.InsertHeader(header)
	}
}