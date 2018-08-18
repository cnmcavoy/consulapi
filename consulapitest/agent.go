// Code autogenerated by mockery v2.0.0
//
// Do not manually edit the content of this file.

// Package consulapitest contains autogenerated mocks.
package consulapitest

import "github.com/shoenig/consulapi"
import "github.com/stretchr/testify/mock"

// Agent is an autogenerated mock type for the Agent type
type Agent struct {
	mock.Mock
}

// MaintenanceMode provides a mock function with given fields: enabled, reason
func (mockerySelf *Agent) MaintenanceMode(enabled bool, reason string) error {
	ret := mockerySelf.Called(enabled, reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, string) error); ok {
		r0 = rf(enabled, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Members provides a mock function with given fields: wan
func (mockerySelf *Agent) Members(wan bool) ([]consulapi.AgentInfo, error) {
	ret := mockerySelf.Called(wan)

	var r0 []consulapi.AgentInfo
	if rf, ok := ret.Get(0).(func(bool) []consulapi.AgentInfo); ok {
		r0 = rf(wan)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]consulapi.AgentInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(wan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reload provides a mock function with given fields:
func (mockerySelf *Agent) Reload() error {
	ret := mockerySelf.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Self provides a mock function with given fields:
func (mockerySelf *Agent) Self() (consulapi.AgentInfo, error) {
	ret := mockerySelf.Called()

	var r0 consulapi.AgentInfo
	if rf, ok := ret.Get(0).(func() consulapi.AgentInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(consulapi.AgentInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}