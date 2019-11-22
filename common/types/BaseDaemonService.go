package types

import (
	"reflect"
	"sync/atomic"
)

type BaseDaemonService struct {
	daemonList CaseList
	handler	[]HandleRequestInterface
	runing int32
	abort chan struct{}
}

func (BD *BaseDaemonService) AddDaemonChannel(ch ChannelInterface) ResponseInterface {
	BD.daemonList.AddReceive(ch.Channel())
	BD.handler = append(BD.handler,ch.Handler())
	return nil
}
func (BD *BaseDaemonService) RemoveDaemonChannel(ch ChannelInterface) ResponseInterface{
	index := BD.daemonList.Find(ch.Channel())
	if index >= 0 {
		BD.daemonList.Delete(index)
		BD.handler = append(BD.handler[:index],BD.handler[index+1:]...)
	}
	return nil
}
func (BD *BaseDaemonService) Start()ResponseInterface{
	if atomic.CompareAndSwapInt32(&BD.runing,0,1) {
		BD.abort = make(chan struct{})
		go BD.loop()
	}else{

	}
	return nil
}
func (BD *BaseDaemonService) loop(){
	for {
		select {
			case chosen, recv, recvOK := reflect.Select(BD.daemonList):
				if recvOK {
					BD.handler[chosen].Handle(recv)
				}
			case <-BD.abort:
				return
		}
	}

}
func (BD *BaseDaemonService) Stop()ResponseInterface{
	if atomic.CompareAndSwapInt32(&BD.runing,1,0) {
		close(BD.abort)
		BD.abort = nil
	}
	return nil
}