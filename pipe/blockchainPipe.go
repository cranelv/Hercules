package pipe

import "github.com/cranelv/Hercules/common/types"

type BlockchainPipe struct {
	pipe PipeInterface
}
func (bcPipe *BlockchainPipe) InsertBlock(block types.BlockInterface,callback func(interface{},error)){
	bcPipe.pipe.Request(Message{BLOCKCHAIN_InsertBlock,block},callback)
}
func (bcPipe *BlockchainPipe) InsertHeader(header types.HeaderInterface,callback func(interface{},error)){
	bcPipe.pipe.Request(Message{BLOCKCHAIN_InsertHeader,header},callback)
}
func (bcPipe *BlockchainPipe) DeleteBlock(block types.BlockInterface,callback func(interface{},error)){
	bcPipe.pipe.Request(Message{BLOCKCHAIN_DeleteBlock,block},callback)
}
