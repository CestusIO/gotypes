// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"code.cestus.io/libs/gotypes/pkg/types"
	"sync"
)

// Ensure, that IDProviderMock does implement types.IDProvider.
// If this is not the case, regenerate this file with moq.
var _ types.IDProvider = &IDProviderMock{}

// IDProviderMock is a mock implementation of types.IDProvider.
//
//	func TestSomethingThatUsesIDProvider(t *testing.T) {
//
//		// make and configure a mocked types.IDProvider
//		mockedIDProvider := &IDProviderMock{
//			FromStringFunc: func(a types.ID, s string) error {
//				panic("mock out the FromString method")
//			},
//			FromStringOrNilFunc: func(a types.ID, s string)  {
//				panic("mock out the FromStringOrNil method")
//			},
//			NewRandomFunc: func(a types.ID) error {
//				panic("mock out the NewRandom method")
//			},
//		}
//
//		// use mockedIDProvider in code that requires types.IDProvider
//		// and then make assertions.
//
//	}
type IDProviderMock struct {
	// FromStringFunc mocks the FromString method.
	FromStringFunc func(a types.ID, s string) error

	// FromStringOrNilFunc mocks the FromStringOrNil method.
	FromStringOrNilFunc func(a types.ID, s string)

	// NewRandomFunc mocks the NewRandom method.
	NewRandomFunc func(a types.ID) error

	// calls tracks calls to the methods.
	calls struct {
		// FromString holds details about calls to the FromString method.
		FromString []struct {
			// A is the a argument value.
			A types.ID
			// S is the s argument value.
			S string
		}
		// FromStringOrNil holds details about calls to the FromStringOrNil method.
		FromStringOrNil []struct {
			// A is the a argument value.
			A types.ID
			// S is the s argument value.
			S string
		}
		// NewRandom holds details about calls to the NewRandom method.
		NewRandom []struct {
			// A is the a argument value.
			A types.ID
		}
	}
	lockFromString      sync.RWMutex
	lockFromStringOrNil sync.RWMutex
	lockNewRandom       sync.RWMutex
}

// FromString calls FromStringFunc.
func (mock *IDProviderMock) FromString(a types.ID, s string) error {
	if mock.FromStringFunc == nil {
		panic("IDProviderMock.FromStringFunc: method is nil but IDProvider.FromString was just called")
	}
	callInfo := struct {
		A types.ID
		S string
	}{
		A: a,
		S: s,
	}
	mock.lockFromString.Lock()
	mock.calls.FromString = append(mock.calls.FromString, callInfo)
	mock.lockFromString.Unlock()
	return mock.FromStringFunc(a, s)
}

// FromStringCalls gets all the calls that were made to FromString.
// Check the length with:
//
//	len(mockedIDProvider.FromStringCalls())
func (mock *IDProviderMock) FromStringCalls() []struct {
	A types.ID
	S string
} {
	var calls []struct {
		A types.ID
		S string
	}
	mock.lockFromString.RLock()
	calls = mock.calls.FromString
	mock.lockFromString.RUnlock()
	return calls
}

// FromStringOrNil calls FromStringOrNilFunc.
func (mock *IDProviderMock) FromStringOrNil(a types.ID, s string) {
	if mock.FromStringOrNilFunc == nil {
		panic("IDProviderMock.FromStringOrNilFunc: method is nil but IDProvider.FromStringOrNil was just called")
	}
	callInfo := struct {
		A types.ID
		S string
	}{
		A: a,
		S: s,
	}
	mock.lockFromStringOrNil.Lock()
	mock.calls.FromStringOrNil = append(mock.calls.FromStringOrNil, callInfo)
	mock.lockFromStringOrNil.Unlock()
	mock.FromStringOrNilFunc(a, s)
}

// FromStringOrNilCalls gets all the calls that were made to FromStringOrNil.
// Check the length with:
//
//	len(mockedIDProvider.FromStringOrNilCalls())
func (mock *IDProviderMock) FromStringOrNilCalls() []struct {
	A types.ID
	S string
} {
	var calls []struct {
		A types.ID
		S string
	}
	mock.lockFromStringOrNil.RLock()
	calls = mock.calls.FromStringOrNil
	mock.lockFromStringOrNil.RUnlock()
	return calls
}

// NewRandom calls NewRandomFunc.
func (mock *IDProviderMock) NewRandom(a types.ID) error {
	if mock.NewRandomFunc == nil {
		panic("IDProviderMock.NewRandomFunc: method is nil but IDProvider.NewRandom was just called")
	}
	callInfo := struct {
		A types.ID
	}{
		A: a,
	}
	mock.lockNewRandom.Lock()
	mock.calls.NewRandom = append(mock.calls.NewRandom, callInfo)
	mock.lockNewRandom.Unlock()
	return mock.NewRandomFunc(a)
}

// NewRandomCalls gets all the calls that were made to NewRandom.
// Check the length with:
//
//	len(mockedIDProvider.NewRandomCalls())
func (mock *IDProviderMock) NewRandomCalls() []struct {
	A types.ID
} {
	var calls []struct {
		A types.ID
	}
	mock.lockNewRandom.RLock()
	calls = mock.calls.NewRandom
	mock.lockNewRandom.RUnlock()
	return calls
}
