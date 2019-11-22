package stateManager

import "github.com/cranelv/Hercules/common/data"

type StateMachineInterface interface {
	GetReadableState(data.Hash)StateSetInterface
	GetWritableState(data.Hash)StateSetInterface
	WriteState(data.Hash,StateSetInterface) error
}
