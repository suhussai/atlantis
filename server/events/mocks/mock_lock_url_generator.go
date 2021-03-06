// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/hootsuite/atlantis/server/events (interfaces: LockURLGenerator)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockLockURLGenerator struct {
	fail func(message string, callerSkip ...int)
}

func NewMockLockURLGenerator() *MockLockURLGenerator {
	return &MockLockURLGenerator{fail: pegomock.GlobalFailHandler}
}

func (mock *MockLockURLGenerator) SetLockURL(_param0 func(string) string) {
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetLockURL", params, []reflect.Type{})
}

func (mock *MockLockURLGenerator) VerifyWasCalledOnce() *VerifierLockURLGenerator {
	return &VerifierLockURLGenerator{mock, pegomock.Times(1), nil}
}

func (mock *MockLockURLGenerator) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierLockURLGenerator {
	return &VerifierLockURLGenerator{mock, invocationCountMatcher, nil}
}

func (mock *MockLockURLGenerator) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierLockURLGenerator {
	return &VerifierLockURLGenerator{mock, invocationCountMatcher, inOrderContext}
}

type VerifierLockURLGenerator struct {
	mock                   *MockLockURLGenerator
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierLockURLGenerator) SetLockURL(_param0 func(string) string) *LockURLGenerator_SetLockURL_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetLockURL", params)
	return &LockURLGenerator_SetLockURL_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type LockURLGenerator_SetLockURL_OngoingVerification struct {
	mock              *MockLockURLGenerator
	methodInvocations []pegomock.MethodInvocation
}

func (c *LockURLGenerator_SetLockURL_OngoingVerification) GetCapturedArguments() func(string) string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *LockURLGenerator_SetLockURL_OngoingVerification) GetAllCapturedArguments() (_param0 []func(string) string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]func(string) string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(func(string) string)
		}
	}
	return
}
