// Copyright 2023 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: antrea.io/antrea/pkg/agent/servicecidr (interfaces: Interface)

// Package testing is a generated GoMock package.
package testing

import (
	servicecidr "antrea.io/antrea/pkg/agent/servicecidr"
	gomock "github.com/golang/mock/gomock"
	net "net"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockInterface) AddEventHandler(arg0 servicecidr.EventHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddEventHandler", arg0)
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockInterfaceMockRecorder) AddEventHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockInterface)(nil).AddEventHandler), arg0)
}

// GetServiceCIDR mocks base method
func (m *MockInterface) GetServiceCIDR(arg0 bool) (*net.IPNet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceCIDR", arg0)
	ret0, _ := ret[0].(*net.IPNet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceCIDR indicates an expected call of GetServiceCIDR
func (mr *MockInterfaceMockRecorder) GetServiceCIDR(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceCIDR", reflect.TypeOf((*MockInterface)(nil).GetServiceCIDR), arg0)
}
