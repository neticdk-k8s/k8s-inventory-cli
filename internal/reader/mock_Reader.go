// Code generated by mockery v2.42.2. DO NOT EDIT.

package reader

import mock "github.com/stretchr/testify/mock"

// MockReader is an autogenerated mock type for the Reader type
type MockReader struct {
	mock.Mock
}

type MockReader_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReader) EXPECT() *MockReader_Expecter {
	return &MockReader_Expecter{mock: &_m.Mock}
}

// ReadString provides a mock function with given fields: prompt
func (_m *MockReader) ReadString(prompt string) (string, error) {
	ret := _m.Called(prompt)

	if len(ret) == 0 {
		panic("no return value specified for ReadString")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(prompt)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(prompt)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(prompt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReader_ReadString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadString'
type MockReader_ReadString_Call struct {
	*mock.Call
}

// ReadString is a helper method to define mock.On call
//   - prompt string
func (_e *MockReader_Expecter) ReadString(prompt interface{}) *MockReader_ReadString_Call {
	return &MockReader_ReadString_Call{Call: _e.mock.On("ReadString", prompt)}
}

func (_c *MockReader_ReadString_Call) Run(run func(prompt string)) *MockReader_ReadString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockReader_ReadString_Call) Return(_a0 string, _a1 error) *MockReader_ReadString_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReader_ReadString_Call) RunAndReturn(run func(string) (string, error)) *MockReader_ReadString_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockReader creates a new instance of MockReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReader {
	mock := &MockReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
