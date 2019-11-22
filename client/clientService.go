package client

import (
	"github.com/cranelv/Hercules/pipe"
	"github.com/cranelv/Hercules/common/types"
	"github.com/cranelv/Hercules/common/protocol"
)
type ClientPipeSet struct {
	txPoolPipe pipe.TxPoolHandlePipe
	rsPipe pipe.RunningSpacePipe

}
func (cs *ClientPipeSet)Connect() error {
	err := cs.txPoolPipe.Connect()
	if err != nil {
		return err
	}
	err = cs.rsPipe.Connect()
	if err != nil {
		return err
	}
	return nil
}
type ClientService struct {
	Pipes ClientPipeSet
}


func (cs *ClientService)SendTx(tx types.TransactionInterface) {
	cs.Pipes.rsPipe.CreateRunningSpaceByNumber(protocol.BlockNumber_Pending,pipe.RunningSpace_Writable,func(result interface{}, err error) {
		if err != nil {
		//			return err
		}
		rsName := result.(string)
		cs.Pipes.rsPipe.TxCall(rsName,tx, func(result interface{}, err error) {
			if err != nil {
				//			return err
			}
			//		result := <-response
			resp := result.(types.ResponseInterface)
			if resp != nil && resp.Error() != nil{
				//			return resp.Error()
			}
			cs.Pipes.txPoolPipe.AppendTx(tx, func(result interface{}, err error) {
				if err != nil {
					//			return err
				}
				//		result := <-response
				resp := result.(types.ResponseInterface)
				if resp != nil && resp.Error() != nil{
					//			return resp.Error()
				}

			})
		})

	})

}

