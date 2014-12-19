// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/squeedee/go_di/widget"
)

type FakeRunner struct {
	RunStub        func() string
	runMutex       sync.RWMutex
	runArgsForCall []struct{}
	runReturns struct {
		result1 string
	}
}

func (fake *FakeRunner) Run() string {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct{}{})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub()
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeRunner) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeRunner) RunReturns(result1 string) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 string
	}{result1}
}

var _ widget.Runner = new(FakeRunner)
