package blockChainService

import "github.com/cranelv/Hercules/common/types"

type InsertBlockHandler struct {
	bkChan chan types.BlockInterface
	blockchain BlockChainInterface
}
func (ins *InsertBlockHandler)Channel() interface{} {
	return ins.bkChan
}
func (ins *InsertBlockHandler)Handler()types.HandleRequestInterface{
	return ins
}

func(ins *InsertBlockHandler)Handle(Data interface{}){
	if Data != nil{
		block := Data.(types.BlockInterface)
		if block == nil{
			return
		}
		ins.blockchain.InsertBlock(block)
	}
}