package pipe

import (
	"github.com/MatrixAINetwork/go-matrix/sharding/protocol/network"
)

type GatewayPipe struct {
	pipe PipeInterface
}
func (gate *GatewayPipe)Connect()error{
	return gate.pipe.Connect()
}
func (gate *GatewayPipe)BroadcastMssage(msg interface{},filter network.PeerFilterInterface,callback func(interface{},error)){
	gate.pipe.Request(Message{GATEWAY_BroadcastMessage,[]interface{}{msg,filter}},callback)
}
