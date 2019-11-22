package blockService

import (
	"github.com/cranelv/Hercules/common/types"
	"github.com/cranelv/Hercules/pipe"
)

type BlockGenerateHandler struct {
	headerChan chan interface{}
	txpoolPipe pipe.TxPoolHandlePipe
	headerPipe pipe.HeaderHandlePipe
	blockHandle BlockHandleInterface
}
func (gh *BlockGenerateHandler)Channel() interface{} {
	return gh.headerChan
}
func (gh *BlockGenerateHandler)Handler()types.HandleRequestInterface{
	return gh
}

func(gh *BlockGenerateHandler)Handle(Data interface{}){
	if Data != nil{
		header := Data.(types.HeaderInterface)
		if header == nil{
			return
		}
		errChan := make(chan error,1)
		gh.txpoolPipe.GenerateBlock(header, func(result interface{}, err error) {
			if err != nil{
				errChan <- err
			}else{
				//result
				txBlock := result.(types.TxBlockInterface)
				gh.headerPipe.GenerateHeader(header,txBlock, func(result interface{}, err error) {
					if err != nil{
						errChan <- err
					}else{
						//result
						header := result.(types.HeaderInterface)
						gh.blockHandle.IntegrateBlock(header,txBlock)
						errChan <- nil
					}
				})
			}
		})
	}
}

