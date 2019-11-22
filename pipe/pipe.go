package pipe
type Message struct {
	code uint64
	data interface{}
}
type PipeInterface interface {
	Name()string
	Connect()error
	Send(Message)error
	Subscribe(Message)error
	Request(Message,func(interface{},error))
}
