// Code generated by counterfeiter. DO NOT EDIT.
package sharedactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/util/clissh"
)

type FakeSecureShellClient struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	ConnectStub        func(string, string, string, string, bool) error
	connectMutex       sync.RWMutex
	connectArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 bool
	}
	connectReturns struct {
		result1 error
	}
	connectReturnsOnCall map[int]struct {
		result1 error
	}
	InteractiveSessionStub        func([]string, clissh.TTYRequest) error
	interactiveSessionMutex       sync.RWMutex
	interactiveSessionArgsForCall []struct {
		arg1 []string
		arg2 clissh.TTYRequest
	}
	interactiveSessionReturns struct {
		result1 error
	}
	interactiveSessionReturnsOnCall map[int]struct {
		result1 error
	}
	LocalPortForwardStub        func([]clissh.LocalPortForward) error
	localPortForwardMutex       sync.RWMutex
	localPortForwardArgsForCall []struct {
		arg1 []clissh.LocalPortForward
	}
	localPortForwardReturns struct {
		result1 error
	}
	localPortForwardReturnsOnCall map[int]struct {
		result1 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
	}
	waitReturns struct {
		result1 error
	}
	waitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecureShellClient) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeSecureShellClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSecureShellClient) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeSecureShellClient) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) Connect(arg1 string, arg2 string, arg3 string, arg4 string, arg5 bool) error {
	fake.connectMutex.Lock()
	ret, specificReturn := fake.connectReturnsOnCall[len(fake.connectArgsForCall)]
	fake.connectArgsForCall = append(fake.connectArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 bool
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Connect", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.connectMutex.Unlock()
	if fake.ConnectStub != nil {
		return fake.ConnectStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.connectReturns
	return fakeReturns.result1
}

func (fake *FakeSecureShellClient) ConnectCallCount() int {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return len(fake.connectArgsForCall)
}

func (fake *FakeSecureShellClient) ConnectCalls(stub func(string, string, string, string, bool) error) {
	fake.connectMutex.Lock()
	defer fake.connectMutex.Unlock()
	fake.ConnectStub = stub
}

func (fake *FakeSecureShellClient) ConnectArgsForCall(i int) (string, string, string, string, bool) {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	argsForCall := fake.connectArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeSecureShellClient) ConnectReturns(result1 error) {
	fake.connectMutex.Lock()
	defer fake.connectMutex.Unlock()
	fake.ConnectStub = nil
	fake.connectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) ConnectReturnsOnCall(i int, result1 error) {
	fake.connectMutex.Lock()
	defer fake.connectMutex.Unlock()
	fake.ConnectStub = nil
	if fake.connectReturnsOnCall == nil {
		fake.connectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.connectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) InteractiveSession(arg1 []string, arg2 clissh.TTYRequest) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.interactiveSessionMutex.Lock()
	ret, specificReturn := fake.interactiveSessionReturnsOnCall[len(fake.interactiveSessionArgsForCall)]
	fake.interactiveSessionArgsForCall = append(fake.interactiveSessionArgsForCall, struct {
		arg1 []string
		arg2 clissh.TTYRequest
	}{arg1Copy, arg2})
	fake.recordInvocation("InteractiveSession", []interface{}{arg1Copy, arg2})
	fake.interactiveSessionMutex.Unlock()
	if fake.InteractiveSessionStub != nil {
		return fake.InteractiveSessionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.interactiveSessionReturns
	return fakeReturns.result1
}

func (fake *FakeSecureShellClient) InteractiveSessionCallCount() int {
	fake.interactiveSessionMutex.RLock()
	defer fake.interactiveSessionMutex.RUnlock()
	return len(fake.interactiveSessionArgsForCall)
}

func (fake *FakeSecureShellClient) InteractiveSessionCalls(stub func([]string, clissh.TTYRequest) error) {
	fake.interactiveSessionMutex.Lock()
	defer fake.interactiveSessionMutex.Unlock()
	fake.InteractiveSessionStub = stub
}

func (fake *FakeSecureShellClient) InteractiveSessionArgsForCall(i int) ([]string, clissh.TTYRequest) {
	fake.interactiveSessionMutex.RLock()
	defer fake.interactiveSessionMutex.RUnlock()
	argsForCall := fake.interactiveSessionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSecureShellClient) InteractiveSessionReturns(result1 error) {
	fake.interactiveSessionMutex.Lock()
	defer fake.interactiveSessionMutex.Unlock()
	fake.InteractiveSessionStub = nil
	fake.interactiveSessionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) InteractiveSessionReturnsOnCall(i int, result1 error) {
	fake.interactiveSessionMutex.Lock()
	defer fake.interactiveSessionMutex.Unlock()
	fake.InteractiveSessionStub = nil
	if fake.interactiveSessionReturnsOnCall == nil {
		fake.interactiveSessionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.interactiveSessionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) LocalPortForward(arg1 []clissh.LocalPortForward) error {
	var arg1Copy []clissh.LocalPortForward
	if arg1 != nil {
		arg1Copy = make([]clissh.LocalPortForward, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.localPortForwardMutex.Lock()
	ret, specificReturn := fake.localPortForwardReturnsOnCall[len(fake.localPortForwardArgsForCall)]
	fake.localPortForwardArgsForCall = append(fake.localPortForwardArgsForCall, struct {
		arg1 []clissh.LocalPortForward
	}{arg1Copy})
	fake.recordInvocation("LocalPortForward", []interface{}{arg1Copy})
	fake.localPortForwardMutex.Unlock()
	if fake.LocalPortForwardStub != nil {
		return fake.LocalPortForwardStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.localPortForwardReturns
	return fakeReturns.result1
}

func (fake *FakeSecureShellClient) LocalPortForwardCallCount() int {
	fake.localPortForwardMutex.RLock()
	defer fake.localPortForwardMutex.RUnlock()
	return len(fake.localPortForwardArgsForCall)
}

func (fake *FakeSecureShellClient) LocalPortForwardCalls(stub func([]clissh.LocalPortForward) error) {
	fake.localPortForwardMutex.Lock()
	defer fake.localPortForwardMutex.Unlock()
	fake.LocalPortForwardStub = stub
}

func (fake *FakeSecureShellClient) LocalPortForwardArgsForCall(i int) []clissh.LocalPortForward {
	fake.localPortForwardMutex.RLock()
	defer fake.localPortForwardMutex.RUnlock()
	argsForCall := fake.localPortForwardArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSecureShellClient) LocalPortForwardReturns(result1 error) {
	fake.localPortForwardMutex.Lock()
	defer fake.localPortForwardMutex.Unlock()
	fake.LocalPortForwardStub = nil
	fake.localPortForwardReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) LocalPortForwardReturnsOnCall(i int, result1 error) {
	fake.localPortForwardMutex.Lock()
	defer fake.localPortForwardMutex.Unlock()
	fake.LocalPortForwardStub = nil
	if fake.localPortForwardReturnsOnCall == nil {
		fake.localPortForwardReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.localPortForwardReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) Wait() error {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
	}{})
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.waitReturns
	return fakeReturns.result1
}

func (fake *FakeSecureShellClient) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeSecureShellClient) WaitCalls(stub func() error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *FakeSecureShellClient) WaitReturns(result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) WaitReturnsOnCall(i int, result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	fake.interactiveSessionMutex.RLock()
	defer fake.interactiveSessionMutex.RUnlock()
	fake.localPortForwardMutex.RLock()
	defer fake.localPortForwardMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecureShellClient) recordInvocation(key string, args []interface{}) {
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

var _ sharedaction.SecureShellClient = new(FakeSecureShellClient)
