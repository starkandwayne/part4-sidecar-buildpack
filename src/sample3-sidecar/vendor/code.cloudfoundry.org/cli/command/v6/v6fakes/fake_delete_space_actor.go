// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeDeleteSpaceActor struct {
	DeleteSpaceByNameAndOrganizationNameStub        func(string, string) (v2action.Warnings, error)
	deleteSpaceByNameAndOrganizationNameMutex       sync.RWMutex
	deleteSpaceByNameAndOrganizationNameArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deleteSpaceByNameAndOrganizationNameReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	deleteSpaceByNameAndOrganizationNameReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteSpaceActor) DeleteSpaceByNameAndOrganizationName(arg1 string, arg2 string) (v2action.Warnings, error) {
	fake.deleteSpaceByNameAndOrganizationNameMutex.Lock()
	ret, specificReturn := fake.deleteSpaceByNameAndOrganizationNameReturnsOnCall[len(fake.deleteSpaceByNameAndOrganizationNameArgsForCall)]
	fake.deleteSpaceByNameAndOrganizationNameArgsForCall = append(fake.deleteSpaceByNameAndOrganizationNameArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DeleteSpaceByNameAndOrganizationName", []interface{}{arg1, arg2})
	fake.deleteSpaceByNameAndOrganizationNameMutex.Unlock()
	if fake.DeleteSpaceByNameAndOrganizationNameStub != nil {
		return fake.DeleteSpaceByNameAndOrganizationNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteSpaceByNameAndOrganizationNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeleteSpaceActor) DeleteSpaceByNameAndOrganizationNameCallCount() int {
	fake.deleteSpaceByNameAndOrganizationNameMutex.RLock()
	defer fake.deleteSpaceByNameAndOrganizationNameMutex.RUnlock()
	return len(fake.deleteSpaceByNameAndOrganizationNameArgsForCall)
}

func (fake *FakeDeleteSpaceActor) DeleteSpaceByNameAndOrganizationNameCalls(stub func(string, string) (v2action.Warnings, error)) {
	fake.deleteSpaceByNameAndOrganizationNameMutex.Lock()
	defer fake.deleteSpaceByNameAndOrganizationNameMutex.Unlock()
	fake.DeleteSpaceByNameAndOrganizationNameStub = stub
}

func (fake *FakeDeleteSpaceActor) DeleteSpaceByNameAndOrganizationNameArgsForCall(i int) (string, string) {
	fake.deleteSpaceByNameAndOrganizationNameMutex.RLock()
	defer fake.deleteSpaceByNameAndOrganizationNameMutex.RUnlock()
	argsForCall := fake.deleteSpaceByNameAndOrganizationNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDeleteSpaceActor) DeleteSpaceByNameAndOrganizationNameReturns(result1 v2action.Warnings, result2 error) {
	fake.deleteSpaceByNameAndOrganizationNameMutex.Lock()
	defer fake.deleteSpaceByNameAndOrganizationNameMutex.Unlock()
	fake.DeleteSpaceByNameAndOrganizationNameStub = nil
	fake.deleteSpaceByNameAndOrganizationNameReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteSpaceActor) DeleteSpaceByNameAndOrganizationNameReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.deleteSpaceByNameAndOrganizationNameMutex.Lock()
	defer fake.deleteSpaceByNameAndOrganizationNameMutex.Unlock()
	fake.DeleteSpaceByNameAndOrganizationNameStub = nil
	if fake.deleteSpaceByNameAndOrganizationNameReturnsOnCall == nil {
		fake.deleteSpaceByNameAndOrganizationNameReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.deleteSpaceByNameAndOrganizationNameReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteSpaceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteSpaceByNameAndOrganizationNameMutex.RLock()
	defer fake.deleteSpaceByNameAndOrganizationNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteSpaceActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.DeleteSpaceActor = new(FakeDeleteSpaceActor)
