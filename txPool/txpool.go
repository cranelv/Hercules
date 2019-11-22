package txPool

import "github.com/cranelv/Hercules/common/types"

type TxPoolInterface interface {
	Append(types.TransactionInterface) error
	Remove(types.TransactionInterface) error
	getPendingTransactions()[]types.TransactionInterface
	SealBlock(string, []types.TransactionInterface)(types.TxBlockInterface,error)
	CommitBlock(types.TxBlockInterface)error
}
