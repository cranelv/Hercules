package blockService

import (
	"github.com/MatrixAINetwork/go-matrix/sharding/types"
)

type BlockService struct {

}
func (bs *BlockService)GenerateBlock(parent types.BlockInterface)error{
	errChan := make(chan error,1)
	bs.txpoolPipe.GenerateBlock(parent.Header(), func(result interface{}, err error) {
		if err != nil{
			errChan <- err
		}else{
			//result
			txBlock := result.(types.TxBlockInterface)
			bs.headerPipe.GenerateHeader(parent.Header(),txBlock, func(result interface{}, err error) {
				if err != nil{
					errChan <- err
				}else{
					//result
					header := result.(types.HeaderInterface)
					bs.IntegrateBlock(header,txBlock)
					errChan <- nil
				}
			})
		}
	})
	err := <- errChan
	if err != nil{
		return err
	}

	return nil
}
func (bs *BlockService)IntegrateBlock(header types.HeaderInterface,txBlock types.TxBlockInterface)(types.BlockInterface,error){
	return nil, nil
}