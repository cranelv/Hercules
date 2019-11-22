package pipe

import (
	"github.com/MatrixAINetwork/go-matrix/sharding/types"
	"github.com/MatrixAINetwork/go-matrix/common"
)
type RunningSpaceStyle int
const (
	RunningSpace_Readable RunningSpaceStyle = iota + 0
	RunningSpace_Writable
)

type RunningSpacePipe struct {
	pipe PipeInterface
}
func (rp *RunningSpacePipe)Connect()error{
	return rp.pipe.Connect()
}
func (rp *RunningSpacePipe)CreateRunningSpaceByNumber(number int64,style RunningSpaceStyle,callback func(interface{},error)){
	rp.pipe.Request(Message{RunningSpace_Create,[]interface{}{number,style,0}},callback)
}
func (rp *RunningSpacePipe)CreateRunningSpaceByHash(hash common.Hash,style RunningSpaceStyle,callback func(interface{},error)){
	rp.pipe.Request(Message{RunningSpace_Create,[]interface{}{hash,style,1}},callback)
}
func (rp *RunningSpacePipe)TxCall(runningSpaceName string,tx types.TransactionInterface,callback func(interface{},error)){
	rp.pipe.Request(Message{RunningSpace_TX_Call,[]interface{}{runningSpaceName,tx}},callback)
}
func (rp *RunningSpacePipe)TxCommit(runningSpaceName string,tx types.TransactionInterface,callback func(interface{},error)){
	rp.pipe.Request(Message{RunningSpace_TX_Commit,[]interface{}{runningSpaceName,tx}},callback)
}
func (rp *RunningSpacePipe)TxLocalCall(runningSpaceName string,tx types.TransactionInterface,callback func(interface{},error)){
	rp.pipe.Request(Message{RunningSpace_TX_LocalCall,[]interface{}{runningSpaceName,tx}},callback)
}