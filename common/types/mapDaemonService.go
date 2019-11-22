package types

import (
	"sync/atomic"
)

type MapDaemonService struct {
	handler	map[string] []HandleRequestInterface
	runing int32
	Request chan RequestInterface
	abort chan struct{}
}

func (md *MapDaemonService) AddDaemonChannel(ch ChannelInterface) ResponseInterface {
	key := ch.Channel().(string)
	md.handler[key] = append(md.handler[key],ch.Handler())
	return nil
}
func (md *MapDaemonService) RemoveDaemonChannel(ch ChannelInterface) ResponseInterface{
	key := ch.Channel().(string)
	if handle,exist := md.handler[key];exist{
		for i := 0; i< len(handle) ; i++  {
			if handle[i] == ch.Handler() {
				if len(handle) == 1{
					delete(md.handler,key)
				}else{
					md.handler[key] = append(handle[:i],handle[i:]...)
				}
				break
			}
		}
	}
	return nil
}
func (md *MapDaemonService) Start()ResponseInterface{
	if atomic.CompareAndSwapInt32(&md.runing,0,1) {
		md.abort = make(chan struct{})
		go md.loop()
	}else{

	}
	return nil
}
func (md *MapDaemonService) loop(){
	for {
		select {
		case req := <-md.Request:
			if handlers,exist := md.handler[req.Channel()];exist{
				for _,handle := range handlers{
					handle.Handle(req)
				}
			}else{
				//log error
			}

		case <-md.abort:
			return
		}
	}

}
func (md *MapDaemonService) Stop()ResponseInterface{
	if atomic.CompareAndSwapInt32(&md.runing,1,0) {
		close(md.abort)
		md.abort = nil
	}
	return nil
}