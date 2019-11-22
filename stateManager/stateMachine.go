package stateManager

import "github.com/MatrixAINetwork/go-matrix/common"

type StateMachineInterface interface {
	GetReadableState(common.Hash)StateSetInterface
	GetWritableState(common.Hash)StateSetInterface
	WriteState(common.Hash,StateSetInterface) error
}
