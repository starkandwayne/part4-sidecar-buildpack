// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeCreateSpaceActor struct {
	CreateSpaceStub        func(string, string, string) (v2action.Space, v2action.Warnings, error)
	createSpaceMutex       sync.RWMutex
	createSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createSpaceReturns struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	createSpaceReturnsOnCall map[int]struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	GrantSpaceDeveloperByUsernameStub        func(string, string) (v2action.Warnings, error)
	grantSpaceDeveloperByUsernameMutex       sync.RWMutex
	grantSpaceDeveloperByUsernameArgsForCall []struct {
		arg1 string
		arg2 string
	}
	grantSpaceDeveloperByUsernameReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	grantSpaceDeveloperByUsernameReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	GrantSpaceManagerByUsernameStub        func(string, string, string) (v2action.Warnings, error)
	grantSpaceManagerByUsernameMutex       sync.RWMutex
	grantSpaceManagerByUsernameArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	grantSpaceManagerByUsernameReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	grantSpaceManagerByUsernameReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreateSpaceActor) CreateSpace(arg1 string, arg2 string, arg3 string) (v2action.Space, v2action.Warnings, error) {
	fake.createSpaceMutex.Lock()
	ret, specificReturn := fake.createSpaceReturnsOnCall[len(fake.createSpaceArgsForCall)]
	fake.createSpaceArgsForCall = append(fake.createSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateSpace", []interface{}{arg1, arg2, arg3})
	fake.createSpaceMutex.Unlock()
	if fake.CreateSpaceStub != nil {
		return fake.CreateSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCreateSpaceActor) CreateSpaceCallCount() int {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	return len(fake.createSpaceArgsForCall)
}

func (fake *FakeCreateSpaceActor) CreateSpaceCalls(stub func(string, string, string) (v2action.Space, v2action.Warnings, error)) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = stub
}

func (fake *FakeCreateSpaceActor) CreateSpaceArgsForCall(i int) (string, string, string) {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	argsForCall := fake.createSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCreateSpaceActor) CreateSpaceReturns(result1 v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = nil
	fake.createSpaceReturns = struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCreateSpaceActor) CreateSpaceReturnsOnCall(i int, result1 v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = nil
	if fake.createSpaceReturnsOnCall == nil {
		fake.createSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.Space
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.createSpaceReturnsOnCall[i] = struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCreateSpaceActor) GrantSpaceDeveloperByUsername(arg1 string, arg2 string) (v2action.Warnings, error) {
	fake.grantSpaceDeveloperByUsernameMutex.Lock()
	ret, specificReturn := fake.grantSpaceDeveloperByUsernameReturnsOnCall[len(fake.grantSpaceDeveloperByUsernameArgsForCall)]
	fake.grantSpaceDeveloperByUsernameArgsForCall = append(fake.grantSpaceDeveloperByUsernameArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GrantSpaceDeveloperByUsername", []interface{}{arg1, arg2})
	fake.grantSpaceDeveloperByUsernameMutex.Unlock()
	if fake.GrantSpaceDeveloperByUsernameStub != nil {
		return fake.GrantSpaceDeveloperByUsernameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.grantSpaceDeveloperByUsernameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreateSpaceActor) GrantSpaceDeveloperByUsernameCallCount() int {
	fake.grantSpaceDeveloperByUsernameMutex.RLock()
	defer fake.grantSpaceDeveloperByUsernameMutex.RUnlock()
	return len(fake.grantSpaceDeveloperByUsernameArgsForCall)
}

func (fake *FakeCreateSpaceActor) GrantSpaceDeveloperByUsernameCalls(stub func(string, string) (v2action.Warnings, error)) {
	fake.grantSpaceDeveloperByUsernameMutex.Lock()
	defer fake.grantSpaceDeveloperByUsernameMutex.Unlock()
	fake.GrantSpaceDeveloperByUsernameStub = stub
}

func (fake *FakeCreateSpaceActor) GrantSpaceDeveloperByUsernameArgsForCall(i int) (string, string) {
	fake.grantSpaceDeveloperByUsernameMutex.RLock()
	defer fake.grantSpaceDeveloperByUsernameMutex.RUnlock()
	argsForCall := fake.grantSpaceDeveloperByUsernameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCreateSpaceActor) GrantSpaceDeveloperByUsernameReturns(result1 v2action.Warnings, result2 error) {
	fake.grantSpaceDeveloperByUsernameMutex.Lock()
	defer fake.grantSpaceDeveloperByUsernameMutex.Unlock()
	fake.GrantSpaceDeveloperByUsernameStub = nil
	fake.grantSpaceDeveloperByUsernameReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateSpaceActor) GrantSpaceDeveloperByUsernameReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.grantSpaceDeveloperByUsernameMutex.Lock()
	defer fake.grantSpaceDeveloperByUsernameMutex.Unlock()
	fake.GrantSpaceDeveloperByUsernameStub = nil
	if fake.grantSpaceDeveloperByUsernameReturnsOnCall == nil {
		fake.grantSpaceDeveloperByUsernameReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.grantSpaceDeveloperByUsernameReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateSpaceActor) GrantSpaceManagerByUsername(arg1 string, arg2 string, arg3 string) (v2action.Warnings, error) {
	fake.grantSpaceManagerByUsernameMutex.Lock()
	ret, specificReturn := fake.grantSpaceManagerByUsernameReturnsOnCall[len(fake.grantSpaceManagerByUsernameArgsForCall)]
	fake.grantSpaceManagerByUsernameArgsForCall = append(fake.grantSpaceManagerByUsernameArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GrantSpaceManagerByUsername", []interface{}{arg1, arg2, arg3})
	fake.grantSpaceManagerByUsernameMutex.Unlock()
	if fake.GrantSpaceManagerByUsernameStub != nil {
		return fake.GrantSpaceManagerByUsernameStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.grantSpaceManagerByUsernameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreateSpaceActor) GrantSpaceManagerByUsernameCallCount() int {
	fake.grantSpaceManagerByUsernameMutex.RLock()
	defer fake.grantSpaceManagerByUsernameMutex.RUnlock()
	return len(fake.grantSpaceManagerByUsernameArgsForCall)
}

func (fake *FakeCreateSpaceActor) GrantSpaceManagerByUsernameCalls(stub func(string, string, string) (v2action.Warnings, error)) {
	fake.grantSpaceManagerByUsernameMutex.Lock()
	defer fake.grantSpaceManagerByUsernameMutex.Unlock()
	fake.GrantSpaceManagerByUsernameStub = stub
}

func (fake *FakeCreateSpaceActor) GrantSpaceManagerByUsernameArgsForCall(i int) (string, string, string) {
	fake.grantSpaceManagerByUsernameMutex.RLock()
	defer fake.grantSpaceManagerByUsernameMutex.RUnlock()
	argsForCall := fake.grantSpaceManagerByUsernameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCreateSpaceActor) GrantSpaceManagerByUsernameReturns(result1 v2action.Warnings, result2 error) {
	fake.grantSpaceManagerByUsernameMutex.Lock()
	defer fake.grantSpaceManagerByUsernameMutex.Unlock()
	fake.GrantSpaceManagerByUsernameStub = nil
	fake.grantSpaceManagerByUsernameReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateSpaceActor) GrantSpaceManagerByUsernameReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.grantSpaceManagerByUsernameMutex.Lock()
	defer fake.grantSpaceManagerByUsernameMutex.Unlock()
	fake.GrantSpaceManagerByUsernameStub = nil
	if fake.grantSpaceManagerByUsernameReturnsOnCall == nil {
		fake.grantSpaceManagerByUsernameReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.grantSpaceManagerByUsernameReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateSpaceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	fake.grantSpaceDeveloperByUsernameMutex.RLock()
	defer fake.grantSpaceDeveloperByUsernameMutex.RUnlock()
	fake.grantSpaceManagerByUsernameMutex.RLock()
	defer fake.grantSpaceManagerByUsernameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreateSpaceActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.CreateSpaceActor = new(FakeCreateSpaceActor)
