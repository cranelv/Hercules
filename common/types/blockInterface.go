package types

import "github.com/MatrixAINetwork/go-matrix/common"

type HeaderInterface interface {
	CheckInterface
	Hash()common.Hash
	Number()uint64
}
type TxBlockInterface interface {
	CheckInterface
}
type BlockInterface interface {
	Header()HeaderInterface
	Parent()common.Hash
}