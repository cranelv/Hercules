package types

import (
	"reflect"
	"errors"
)
var errBadRececiveChan = errors.New("CaseList: CaseList Add parameter does not have Receiveable channel type")
var errBadSendChan = errors.New("CaseList: CaseList Add parameter does not have Sendable channel type")

type CaseList []reflect.SelectCase

// find returns the index of a case containing the given channel.
func (cs CaseList) Find(channel interface{}) int {
	for i, cas := range cs {
		if cas.Chan.Interface() == channel {
			return i
		}
	}
	return -1
}

// delete removes the given case from cs.
func (cs CaseList) Delete(index int) CaseList {
	return append(cs[:index], cs[index+1:]...)
}

// deactivate moves the case at index into the non-accessible portion of the cs slice.
func (cs CaseList) Deactivate(index int) CaseList {
	last := len(cs) - 1
	cs[index], cs[last] = cs[last], cs[index]
	return cs[:last]
}
func (cs *CaseList) AddSend(channel interface{}){
	chanval := reflect.ValueOf(channel)
	chantyp := chanval.Type()
	if chantyp.Kind() != reflect.Chan || chantyp.ChanDir()&reflect.SendDir == 0 {
		panic(errBadSendChan)
	}
	// Add the select case to the caseList.
	cas := reflect.SelectCase{Dir: reflect.SelectSend, Chan: chanval}
	*cs = append(*cs, cas)
}
func (cs *CaseList) AddReceive(channel interface{}){
	chanval := reflect.ValueOf(channel)
	chantyp := chanval.Type()
	if chantyp.Kind() != reflect.Chan || chantyp.ChanDir()&reflect.RecvDir == 0 {
		panic(errBadRececiveChan)
	}
	// Add the select case to the caseList.
	cas := reflect.SelectCase{Dir: reflect.SelectSend, Chan: chanval}
	*cs = append(*cs, cas)
}
func (cs *CaseList) AddAbort(abort interface{}){
	chanval := reflect.ValueOf(abort)
	chantyp := chanval.Type()
	if chantyp.Kind() != reflect.Chan{
		panic(errBadSendChan)
	}
	// Add the select case to the caseList.
	cas := reflect.SelectCase{Dir: reflect.SelectSend, Chan: chanval}
	*cs = append(*cs, cas)
}