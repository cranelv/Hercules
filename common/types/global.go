package types

import "github.com/cranelv/Hercules/common/data"

type RequestInterface interface {
	Data() []byte
	Channel() string
}
type ResponseInterface interface {
	Error() error
}
type HandleRequestInterface interface {
	Handle(interface{})
}
type CheckInterface interface {
	Check(RequestInterface) ResponseInterface
}
type TransactionInterface interface {
	From() data.Address
	Catagrate() []byte
}

type TxHandleInterface interface {
	Decode(RequestInterface) (TransactionInterface,error)
	Call(TransactionInterface) ResponseInterface
	Commit(TransactionInterface) ResponseInterface
}
type ChannelInterface interface {
	Channel()interface{}
	Handler()HandleRequestInterface
}
type DaemonServiceInterface interface {
	AddDaemonChannel(ChannelInterface) ResponseInterface
	RemoveDaemonChannel(ChannelInterface) ResponseInterface
	Start()ResponseInterface
	Stop()ResponseInterface
}