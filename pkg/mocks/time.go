// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"code.cestus.io/libs/gotypes/pkg/types"
	"sync"
	"time"
)

// Ensure, that TimeProviderMock does implement types.TimeProvider.
// If this is not the case, regenerate this file with moq.
var _ types.TimeProvider = &TimeProviderMock{}

// TimeProviderMock is a mock implementation of types.TimeProvider.
//
//	func TestSomethingThatUsesTimeProvider(t *testing.T) {
//
//		// make and configure a mocked types.TimeProvider
//		mockedTimeProvider := &TimeProviderMock{
//			NowFunc: func() time.Time {
//				panic("mock out the Now method")
//			},
//		}
//
//		// use mockedTimeProvider in code that requires types.TimeProvider
//		// and then make assertions.
//
//	}
type TimeProviderMock struct {
	// NowFunc mocks the Now method.
	NowFunc func() time.Time

	// calls tracks calls to the methods.
	calls struct {
		// Now holds details about calls to the Now method.
		Now []struct {
		}
	}
	lockNow sync.RWMutex
}

// Now calls NowFunc.
func (mock *TimeProviderMock) Now() time.Time {
	if mock.NowFunc == nil {
		panic("TimeProviderMock.NowFunc: method is nil but TimeProvider.Now was just called")
	}
	callInfo := struct {
	}{}
	mock.lockNow.Lock()
	mock.calls.Now = append(mock.calls.Now, callInfo)
	mock.lockNow.Unlock()
	return mock.NowFunc()
}

// NowCalls gets all the calls that were made to Now.
// Check the length with:
//
//	len(mockedTimeProvider.NowCalls())
func (mock *TimeProviderMock) NowCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockNow.RLock()
	calls = mock.calls.Now
	mock.lockNow.RUnlock()
	return calls
}
