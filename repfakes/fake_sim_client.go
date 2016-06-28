// This file was generated by counterfeiter
package repfakes

import (
	"net/http"
	"sync"
	"time"

	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/rep"
)

type FakeSimClient struct {
	StateStub        func() (rep.CellState, error)
	stateMutex       sync.RWMutex
	stateArgsForCall []struct{}
	stateReturns     struct {
		result1 rep.CellState
		result2 error
	}
	PerformStub        func(work rep.Work) (rep.Work, error)
	performMutex       sync.RWMutex
	performArgsForCall []struct {
		work rep.Work
	}
	performReturns struct {
		result1 rep.Work
		result2 error
	}
	StopLRPInstanceStub        func(key models.ActualLRPKey, instanceKey models.ActualLRPInstanceKey) error
	stopLRPInstanceMutex       sync.RWMutex
	stopLRPInstanceArgsForCall []struct {
		key         models.ActualLRPKey
		instanceKey models.ActualLRPInstanceKey
	}
	stopLRPInstanceReturns struct {
		result1 error
	}
	CancelTaskStub        func(taskGuid string) error
	cancelTaskMutex       sync.RWMutex
	cancelTaskArgsForCall []struct {
		taskGuid string
	}
	cancelTaskReturns struct {
		result1 error
	}
	SetStateClientStub        func(stateClient *http.Client)
	setStateClientMutex       sync.RWMutex
	setStateClientArgsForCall []struct {
		stateClient *http.Client
	}
	StateClientTimeoutStub        func() time.Duration
	stateClientTimeoutMutex       sync.RWMutex
	stateClientTimeoutArgsForCall []struct{}
	stateClientTimeoutReturns     struct {
		result1 time.Duration
	}
	ResetStub        func() error
	resetMutex       sync.RWMutex
	resetArgsForCall []struct{}
	resetReturns     struct {
		result1 error
	}
}

func (fake *FakeSimClient) State() (rep.CellState, error) {
	fake.stateMutex.Lock()
	fake.stateArgsForCall = append(fake.stateArgsForCall, struct{}{})
	fake.stateMutex.Unlock()
	if fake.StateStub != nil {
		return fake.StateStub()
	} else {
		return fake.stateReturns.result1, fake.stateReturns.result2
	}
}

func (fake *FakeSimClient) StateCallCount() int {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return len(fake.stateArgsForCall)
}

func (fake *FakeSimClient) StateReturns(result1 rep.CellState, result2 error) {
	fake.StateStub = nil
	fake.stateReturns = struct {
		result1 rep.CellState
		result2 error
	}{result1, result2}
}

func (fake *FakeSimClient) Perform(work rep.Work) (rep.Work, error) {
	fake.performMutex.Lock()
	fake.performArgsForCall = append(fake.performArgsForCall, struct {
		work rep.Work
	}{work})
	fake.performMutex.Unlock()
	if fake.PerformStub != nil {
		return fake.PerformStub(work)
	} else {
		return fake.performReturns.result1, fake.performReturns.result2
	}
}

func (fake *FakeSimClient) PerformCallCount() int {
	fake.performMutex.RLock()
	defer fake.performMutex.RUnlock()
	return len(fake.performArgsForCall)
}

func (fake *FakeSimClient) PerformArgsForCall(i int) rep.Work {
	fake.performMutex.RLock()
	defer fake.performMutex.RUnlock()
	return fake.performArgsForCall[i].work
}

func (fake *FakeSimClient) PerformReturns(result1 rep.Work, result2 error) {
	fake.PerformStub = nil
	fake.performReturns = struct {
		result1 rep.Work
		result2 error
	}{result1, result2}
}

func (fake *FakeSimClient) StopLRPInstance(key models.ActualLRPKey, instanceKey models.ActualLRPInstanceKey) error {
	fake.stopLRPInstanceMutex.Lock()
	fake.stopLRPInstanceArgsForCall = append(fake.stopLRPInstanceArgsForCall, struct {
		key         models.ActualLRPKey
		instanceKey models.ActualLRPInstanceKey
	}{key, instanceKey})
	fake.stopLRPInstanceMutex.Unlock()
	if fake.StopLRPInstanceStub != nil {
		return fake.StopLRPInstanceStub(key, instanceKey)
	} else {
		return fake.stopLRPInstanceReturns.result1
	}
}

func (fake *FakeSimClient) StopLRPInstanceCallCount() int {
	fake.stopLRPInstanceMutex.RLock()
	defer fake.stopLRPInstanceMutex.RUnlock()
	return len(fake.stopLRPInstanceArgsForCall)
}

func (fake *FakeSimClient) StopLRPInstanceArgsForCall(i int) (models.ActualLRPKey, models.ActualLRPInstanceKey) {
	fake.stopLRPInstanceMutex.RLock()
	defer fake.stopLRPInstanceMutex.RUnlock()
	return fake.stopLRPInstanceArgsForCall[i].key, fake.stopLRPInstanceArgsForCall[i].instanceKey
}

func (fake *FakeSimClient) StopLRPInstanceReturns(result1 error) {
	fake.StopLRPInstanceStub = nil
	fake.stopLRPInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSimClient) CancelTask(taskGuid string) error {
	fake.cancelTaskMutex.Lock()
	fake.cancelTaskArgsForCall = append(fake.cancelTaskArgsForCall, struct {
		taskGuid string
	}{taskGuid})
	fake.cancelTaskMutex.Unlock()
	if fake.CancelTaskStub != nil {
		return fake.CancelTaskStub(taskGuid)
	} else {
		return fake.cancelTaskReturns.result1
	}
}

func (fake *FakeSimClient) CancelTaskCallCount() int {
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	return len(fake.cancelTaskArgsForCall)
}

func (fake *FakeSimClient) CancelTaskArgsForCall(i int) string {
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	return fake.cancelTaskArgsForCall[i].taskGuid
}

func (fake *FakeSimClient) CancelTaskReturns(result1 error) {
	fake.CancelTaskStub = nil
	fake.cancelTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSimClient) SetStateClient(stateClient *http.Client) {
	fake.setStateClientMutex.Lock()
	fake.setStateClientArgsForCall = append(fake.setStateClientArgsForCall, struct {
		stateClient *http.Client
	}{stateClient})
	fake.setStateClientMutex.Unlock()
	if fake.SetStateClientStub != nil {
		fake.SetStateClientStub(stateClient)
	}
}

func (fake *FakeSimClient) SetStateClientCallCount() int {
	fake.setStateClientMutex.RLock()
	defer fake.setStateClientMutex.RUnlock()
	return len(fake.setStateClientArgsForCall)
}

func (fake *FakeSimClient) SetStateClientArgsForCall(i int) *http.Client {
	fake.setStateClientMutex.RLock()
	defer fake.setStateClientMutex.RUnlock()
	return fake.setStateClientArgsForCall[i].stateClient
}

func (fake *FakeSimClient) StateClientTimeout() time.Duration {
	fake.stateClientTimeoutMutex.Lock()
	fake.stateClientTimeoutArgsForCall = append(fake.stateClientTimeoutArgsForCall, struct{}{})
	fake.stateClientTimeoutMutex.Unlock()
	if fake.StateClientTimeoutStub != nil {
		return fake.StateClientTimeoutStub()
	} else {
		return fake.stateClientTimeoutReturns.result1
	}
}

func (fake *FakeSimClient) StateClientTimeoutCallCount() int {
	fake.stateClientTimeoutMutex.RLock()
	defer fake.stateClientTimeoutMutex.RUnlock()
	return len(fake.stateClientTimeoutArgsForCall)
}

func (fake *FakeSimClient) StateClientTimeoutReturns(result1 time.Duration) {
	fake.StateClientTimeoutStub = nil
	fake.stateClientTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeSimClient) Reset() error {
	fake.resetMutex.Lock()
	fake.resetArgsForCall = append(fake.resetArgsForCall, struct{}{})
	fake.resetMutex.Unlock()
	if fake.ResetStub != nil {
		return fake.ResetStub()
	} else {
		return fake.resetReturns.result1
	}
}

func (fake *FakeSimClient) ResetCallCount() int {
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	return len(fake.resetArgsForCall)
}

func (fake *FakeSimClient) ResetReturns(result1 error) {
	fake.ResetStub = nil
	fake.resetReturns = struct {
		result1 error
	}{result1}
}

var _ rep.SimClient = new(FakeSimClient)
