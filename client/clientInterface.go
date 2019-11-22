package client

import "github.com/cranelv/Hercules/common/types"

type ClientInterface interface {
	SendTx(tx types.TransactionInterface)
}
