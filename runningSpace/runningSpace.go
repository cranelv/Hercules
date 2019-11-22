package runningSpace

import (
	"github.com/MatrixAINetwork/go-matrix/sharding/stateManager"
	"github.com/MatrixAINetwork/go-matrix/sharding/types"
)

type RunningSpaceInterface interface {
	Name() string
	StateSet() stateManager.StateSetInterface
	CallTransactions([]types.TransactionInterface)
}