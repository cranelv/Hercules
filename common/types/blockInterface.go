package types

import "github.com/cranelv/Hercules/common/data"

type HeaderInterface interface {
	CheckInterface
	Hash()data.Hash
	Number()uint64
}
type TxBlockInterface interface {
	CheckInterface
}
type BlockInterface interface {
	Header()HeaderInterface
	Parent()data.Hash
}