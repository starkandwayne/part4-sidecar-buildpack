// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeSetHealthCheckActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct {
	}
	cloudControllerAPIVersionReturns struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	SetApplicationProcessHealthCheckTypeByNameAndSpaceStub        func(string, string, constant.HealthCheckType, string, string, int64) (v7action.Application, v7action.Warnings, error)
	setApplicationProcessHealthCheckTypeByNameAndSpaceMutex       sync.RWMutex
	setApplicationProcessHealthCheckTypeByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 constant.HealthCheckType
		arg4 string
		arg5 string
		arg6 int64
	}
	setApplicationProcessHealthCheckTypeByNameAndSpaceReturns struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	setApplicationProcessHealthCheckTypeByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSetHealthCheckActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudControllerAPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeSetHealthCheckActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeSetHealthCheckActor) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeSetHealthCheckActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSetHealthCheckActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSetHealthCheckActor) SetApplicationProcessHealthCheckTypeByNameAndSpace(arg1 string, arg2 string, arg3 constant.HealthCheckType, arg4 string, arg5 string, arg6 int64) (v7action.Application, v7action.Warnings, error) {
	fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.setApplicationProcessHealthCheckTypeByNameAndSpaceReturnsOnCall[len(fake.setApplicationProcessHealthCheckTypeByNameAndSpaceArgsForCall)]
	fake.setApplicationProcessHealthCheckTypeByNameAndSpaceArgsForCall = append(fake.setApplicationProcessHealthCheckTypeByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 constant.HealthCheckType
		arg4 string
		arg5 string
		arg6 int64
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("SetApplicationProcessHealthCheckTypeByNameAndSpace", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.Unlock()
	if fake.SetApplicationProcessHealthCheckTypeByNameAndSpaceStub != nil {
		return fake.SetApplicationProcessHealthCheckTypeByNameAndSpaceStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.setApplicationProcessHealthCheckTypeByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSetHealthCheckActor) SetApplicationProcessHealthCheckTypeByNameAndSpaceCallCount() int {
	fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.RLock()
	defer fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.RUnlock()
	return len(fake.setApplicationProcessHealthCheckTypeByNameAndSpaceArgsForCall)
}

func (fake *FakeSetHealthCheckActor) SetApplicationProcessHealthCheckTypeByNameAndSpaceCalls(stub func(string, string, constant.HealthCheckType, string, string, int64) (v7action.Application, v7action.Warnings, error)) {
	fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.Lock()
	defer fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.Unlock()
	fake.SetApplicationProcessHealthCheckTypeByNameAndSpaceStub = stub
}

func (fake *FakeSetHealthCheckActor) SetApplicationProcessHealthCheckTypeByNameAndSpaceArgsForCall(i int) (string, string, constant.HealthCheckType, string, string, int64) {
	fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.RLock()
	defer fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.setApplicationProcessHealthCheckTypeByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeSetHealthCheckActor) SetApplicationProcessHealthCheckTypeByNameAndSpaceReturns(result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.Lock()
	defer fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.Unlock()
	fake.SetApplicationProcessHealthCheckTypeByNameAndSpaceStub = nil
	fake.setApplicationProcessHealthCheckTypeByNameAndSpaceReturns = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetHealthCheckActor) SetApplicationProcessHealthCheckTypeByNameAndSpaceReturnsOnCall(i int, result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.Lock()
	defer fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.Unlock()
	fake.SetApplicationProcessHealthCheckTypeByNameAndSpaceStub = nil
	if fake.setApplicationProcessHealthCheckTypeByNameAndSpaceReturnsOnCall == nil {
		fake.setApplicationProcessHealthCheckTypeByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.Application
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.setApplicationProcessHealthCheckTypeByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetHealthCheckActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.RLock()
	defer fake.setApplicationProcessHealthCheckTypeByNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSetHealthCheckActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.SetHealthCheckActor = new(FakeSetHealthCheckActor)
