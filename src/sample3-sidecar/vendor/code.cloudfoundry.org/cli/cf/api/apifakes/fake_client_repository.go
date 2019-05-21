// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api"
)

type FakeClientRepository struct {
	ClientExistsStub        func(string) (bool, error)
	clientExistsMutex       sync.RWMutex
	clientExistsArgsForCall []struct {
		arg1 string
	}
	clientExistsReturns struct {
		result1 bool
		result2 error
	}
	clientExistsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClientRepository) ClientExists(arg1 string) (bool, error) {
	fake.clientExistsMutex.Lock()
	ret, specificReturn := fake.clientExistsReturnsOnCall[len(fake.clientExistsArgsForCall)]
	fake.clientExistsArgsForCall = append(fake.clientExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ClientExists", []interface{}{arg1})
	fake.clientExistsMutex.Unlock()
	if fake.ClientExistsStub != nil {
		return fake.ClientExistsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.clientExistsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClientRepository) ClientExistsCallCount() int {
	fake.clientExistsMutex.RLock()
	defer fake.clientExistsMutex.RUnlock()
	return len(fake.clientExistsArgsForCall)
}

func (fake *FakeClientRepository) ClientExistsCalls(stub func(string) (bool, error)) {
	fake.clientExistsMutex.Lock()
	defer fake.clientExistsMutex.Unlock()
	fake.ClientExistsStub = stub
}

func (fake *FakeClientRepository) ClientExistsArgsForCall(i int) string {
	fake.clientExistsMutex.RLock()
	defer fake.clientExistsMutex.RUnlock()
	argsForCall := fake.clientExistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClientRepository) ClientExistsReturns(result1 bool, result2 error) {
	fake.clientExistsMutex.Lock()
	defer fake.clientExistsMutex.Unlock()
	fake.ClientExistsStub = nil
	fake.clientExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeClientRepository) ClientExistsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.clientExistsMutex.Lock()
	defer fake.clientExistsMutex.Unlock()
	fake.ClientExistsStub = nil
	if fake.clientExistsReturnsOnCall == nil {
		fake.clientExistsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.clientExistsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeClientRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clientExistsMutex.RLock()
	defer fake.clientExistsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClientRepository) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ api.ClientRepository = new(FakeClientRepository)
