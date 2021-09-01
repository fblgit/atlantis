// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: WorkingDirLocker)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockWorkingDirLocker struct {
	fail func(message string, callerSkip ...int)
}

func NewMockWorkingDirLocker(options ...pegomock.Option) *MockWorkingDirLocker {
	mock := &MockWorkingDirLocker{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockWorkingDirLocker) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockWorkingDirLocker) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockWorkingDirLocker) TryLock(repoFullName string, pullNum int, workspace string, projectName string) (func(), error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockWorkingDirLocker().")
	}
	params := []pegomock.Param{repoFullName, pullNum, workspace, projectName}
	result := pegomock.GetGenericMockFrom(mock).Invoke("TryLock", params, []reflect.Type{reflect.TypeOf((*func())(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 func()
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(func())
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockWorkingDirLocker) TryLockPull(repoFullName string, pullNum int) (func(), error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockWorkingDirLocker().")
	}
	params := []pegomock.Param{repoFullName, pullNum}
	result := pegomock.GetGenericMockFrom(mock).Invoke("TryLockPull", params, []reflect.Type{reflect.TypeOf((*func())(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 func()
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(func())
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockWorkingDirLocker) VerifyWasCalledOnce() *VerifierMockWorkingDirLocker {
	return &VerifierMockWorkingDirLocker{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockWorkingDirLocker) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockWorkingDirLocker {
	return &VerifierMockWorkingDirLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockWorkingDirLocker) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockWorkingDirLocker {
	return &VerifierMockWorkingDirLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockWorkingDirLocker) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockWorkingDirLocker {
	return &VerifierMockWorkingDirLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockWorkingDirLocker struct {
	mock                   *MockWorkingDirLocker
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockWorkingDirLocker) TryLock(repoFullName string, pullNum int, workspace string, projectName string) *MockWorkingDirLocker_TryLock_OngoingVerification {
	params := []pegomock.Param{repoFullName, pullNum, workspace, projectName}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "TryLock", params, verifier.timeout)
	return &MockWorkingDirLocker_TryLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockWorkingDirLocker_TryLock_OngoingVerification struct {
	mock              *MockWorkingDirLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockWorkingDirLocker_TryLock_OngoingVerification) GetCapturedArguments() (string, int, string, string) {
	repoFullName, pullNum, workspace, projectName := c.GetAllCapturedArguments()
	return repoFullName[len(repoFullName)-1], pullNum[len(pullNum)-1], workspace[len(workspace)-1], projectName[len(projectName)-1]
}

func (c *MockWorkingDirLocker_TryLock_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []int, _param2 []string, _param3 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockWorkingDirLocker) TryLockPull(repoFullName string, pullNum int) *MockWorkingDirLocker_TryLockPull_OngoingVerification {
	params := []pegomock.Param{repoFullName, pullNum}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "TryLockPull", params, verifier.timeout)
	return &MockWorkingDirLocker_TryLockPull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockWorkingDirLocker_TryLockPull_OngoingVerification struct {
	mock              *MockWorkingDirLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockWorkingDirLocker_TryLockPull_OngoingVerification) GetCapturedArguments() (string, int) {
	repoFullName, pullNum := c.GetAllCapturedArguments()
	return repoFullName[len(repoFullName)-1], pullNum[len(pullNum)-1]
}

func (c *MockWorkingDirLocker_TryLockPull_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
	}
	return
}
