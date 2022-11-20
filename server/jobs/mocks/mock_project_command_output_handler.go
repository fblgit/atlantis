// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/jobs (interfaces: ProjectCommandOutputHandler)

package mocks

import (
	"reflect"
	"time"

	pegomock "github.com/petergtz/pegomock"
	"github.com/runatlantis/atlantis/server/events/command"
	jobs "github.com/runatlantis/atlantis/server/jobs"
)

type MockProjectCommandOutputHandler struct {
	fail func(message string, callerSkip ...int)
}

func NewMockProjectCommandOutputHandler(options ...pegomock.Option) *MockProjectCommandOutputHandler {
	mock := &MockProjectCommandOutputHandler{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockProjectCommandOutputHandler) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockProjectCommandOutputHandler) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockProjectCommandOutputHandler) CleanUp(_param0 jobs.PullInfo) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandOutputHandler().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("CleanUp", params, []reflect.Type{})
}

func (mock *MockProjectCommandOutputHandler) Deregister(_param0 string, _param1 chan string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandOutputHandler().")
	}
	params := []pegomock.Param{_param0, _param1}
	pegomock.GetGenericMockFrom(mock).Invoke("Deregister", params, []reflect.Type{})
}

func (mock *MockProjectCommandOutputHandler) Handle() {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandOutputHandler().")
	}
	params := []pegomock.Param{}
	pegomock.GetGenericMockFrom(mock).Invoke("Handle", params, []reflect.Type{})
}

func (mock *MockProjectCommandOutputHandler) IsKeyExists(_param0 string) bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandOutputHandler().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("IsKeyExists", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockProjectCommandOutputHandler) Register(_param0 string, _param1 chan string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandOutputHandler().")
	}
	params := []pegomock.Param{_param0, _param1}
	pegomock.GetGenericMockFrom(mock).Invoke("Register", params, []reflect.Type{})
}

func (mock *MockProjectCommandOutputHandler) Send(_param0 command.ProjectContext, _param1 string, _param2 bool) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandOutputHandler().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	pegomock.GetGenericMockFrom(mock).Invoke("Send", params, []reflect.Type{})
}

func (mock *MockProjectCommandOutputHandler) VerifyWasCalledOnce() *VerifierMockProjectCommandOutputHandler {
	return &VerifierMockProjectCommandOutputHandler{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockProjectCommandOutputHandler) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockProjectCommandOutputHandler {
	return &VerifierMockProjectCommandOutputHandler{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockProjectCommandOutputHandler) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockProjectCommandOutputHandler {
	return &VerifierMockProjectCommandOutputHandler{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockProjectCommandOutputHandler) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockProjectCommandOutputHandler {
	return &VerifierMockProjectCommandOutputHandler{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockProjectCommandOutputHandler struct {
	mock                   *MockProjectCommandOutputHandler
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockProjectCommandOutputHandler) CleanUp(_param0 jobs.PullInfo) *MockProjectCommandOutputHandler_CleanUp_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CleanUp", params, verifier.timeout)
	return &MockProjectCommandOutputHandler_CleanUp_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandOutputHandler_CleanUp_OngoingVerification struct {
	mock              *MockProjectCommandOutputHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandOutputHandler_CleanUp_OngoingVerification) GetCapturedArguments() jobs.PullInfo {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockProjectCommandOutputHandler_CleanUp_OngoingVerification) GetAllCapturedArguments() (_param0 []jobs.PullInfo) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]jobs.PullInfo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(jobs.PullInfo)
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandOutputHandler) Deregister(_param0 string, _param1 chan string) *MockProjectCommandOutputHandler_Deregister_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Deregister", params, verifier.timeout)
	return &MockProjectCommandOutputHandler_Deregister_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandOutputHandler_Deregister_OngoingVerification struct {
	mock              *MockProjectCommandOutputHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandOutputHandler_Deregister_OngoingVerification) GetCapturedArguments() (string, chan string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockProjectCommandOutputHandler_Deregister_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []chan string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]chan string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(chan string)
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandOutputHandler) Handle() *MockProjectCommandOutputHandler_Handle_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Handle", params, verifier.timeout)
	return &MockProjectCommandOutputHandler_Handle_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandOutputHandler_Handle_OngoingVerification struct {
	mock              *MockProjectCommandOutputHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandOutputHandler_Handle_OngoingVerification) GetCapturedArguments() {
}

func (c *MockProjectCommandOutputHandler_Handle_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockProjectCommandOutputHandler) IsKeyExists(_param0 string) *MockProjectCommandOutputHandler_IsKeyExists_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "IsKeyExists", params, verifier.timeout)
	return &MockProjectCommandOutputHandler_IsKeyExists_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandOutputHandler_IsKeyExists_OngoingVerification struct {
	mock              *MockProjectCommandOutputHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandOutputHandler_IsKeyExists_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockProjectCommandOutputHandler_IsKeyExists_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandOutputHandler) Register(_param0 string, _param1 chan string) *MockProjectCommandOutputHandler_Register_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Register", params, verifier.timeout)
	return &MockProjectCommandOutputHandler_Register_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandOutputHandler_Register_OngoingVerification struct {
	mock              *MockProjectCommandOutputHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandOutputHandler_Register_OngoingVerification) GetCapturedArguments() (string, chan string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockProjectCommandOutputHandler_Register_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []chan string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]chan string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(chan string)
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandOutputHandler) Send(_param0 command.ProjectContext, _param1 string, _param2 bool) *MockProjectCommandOutputHandler_Send_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Send", params, verifier.timeout)
	return &MockProjectCommandOutputHandler_Send_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandOutputHandler_Send_OngoingVerification struct {
	mock              *MockProjectCommandOutputHandler
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandOutputHandler_Send_OngoingVerification) GetCapturedArguments() (command.ProjectContext, string, bool) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockProjectCommandOutputHandler_Send_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext, _param1 []string, _param2 []bool) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]command.ProjectContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(command.ProjectContext)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]bool, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(bool)
		}
	}
	return
}
