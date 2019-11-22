package txPool

import (
	"github.com/MatrixAINetwork/go-matrix/sharding/types"
)


type TxPoolService struct {
	types.BaseDaemonService
}

func (ts *TxPoolService)Init() error {
	receive := &ReceiveTxHandler{}
	rep := ts.AddDaemonChannel(receive)
	if rep != nil && rep.Error() != nil{
		return rep.Error()
	}
	return nil
}
