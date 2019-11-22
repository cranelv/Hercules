package runningSpace

import (
	"github.com/cranelv/Hercules/stateManager"
	"github.com/cranelv/Hercules/common/types"
)

type RunningSpaceInterface interface {
	Name() string
	StateSet() stateManager.StateSetInterface
	CallTransactions([]types.TransactionInterface)
}