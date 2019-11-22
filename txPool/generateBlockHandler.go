package txPool

import (
	"github.com/cranelv/Hercules/pipe"
	"github.com/cranelv/Hercules/common/types"
	"sync"
)

type GenerateBlockHandler struct {
	headerChan chan interface{}
	rsPipe pipe.RunningSpacePipe
	TxPoolHandle TxPoolInterface
}
func (gh *GenerateBlockHandler)Channel() interface{} {
	return gh.headerChan
}
func (gh *GenerateBlockHandler)Handler()types.HandleRequestInterface{
	return gh
}

func(gh *GenerateBlockHandler)Handle(Data interface{}){
	if Data != nil{
		header := Data.(types.HeaderInterface)
		if header != nil{
			txs := gh.TxPoolHandle.getPendingTransactions()

			gh.rsPipe.CreateRunningSpaceByHash(header.Hash(),pipe.RunningSpace_Writable, func(result interface{}, err error) {
				if err != nil {
					//return error
				}
				rsName := result.(string)
				collection := make([]types.TransactionInterface,0,len(txs))
				var wg sync.WaitGroup
				wg.Add(len(txs))
				for i:=0;i<len(txs);i++{
					gh.rsPipe.TxCall(rsName,txs[i], func(result interface{}, err error) {
						if err != nil {

						}else{
							collection = append(collection,txs[i])
						}
						wg.Done()
					})
				}
				wg.Wait()
				gh.TxPoolHandle.SealBlock(rsName,collection)

			})

		}
	}
}

