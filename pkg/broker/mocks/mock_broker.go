// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubernetes-sigs/minibroker/pkg/broker (interfaces: MinibrokerClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v2 "github.com/pmorie/go-open-service-broker-client/v2"
	reflect "reflect"
)

// MockMinibrokerClient is a mock of MinibrokerClient interface
type MockMinibrokerClient struct {
	ctrl     *gomock.Controller
	recorder *MockMinibrokerClientMockRecorder
}

// MockMinibrokerClientMockRecorder is the mock recorder for MockMinibrokerClient
type MockMinibrokerClientMockRecorder struct {
	mock *MockMinibrokerClient
}

// NewMockMinibrokerClient creates a new mock instance
func NewMockMinibrokerClient(ctrl *gomock.Controller) *MockMinibrokerClient {
	mock := &MockMinibrokerClient{ctrl: ctrl}
	mock.recorder = &MockMinibrokerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMinibrokerClient) EXPECT() *MockMinibrokerClientMockRecorder {
	return m.recorder
}

// Bind mocks base method
func (m *MockMinibrokerClient) Bind(arg0, arg1, arg2 string, arg3 bool, arg4 map[string]interface{}) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bind indicates an expected call of Bind
func (mr *MockMinibrokerClientMockRecorder) Bind(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockMinibrokerClient)(nil).Bind), arg0, arg1, arg2, arg3, arg4)
}

// Deprovision mocks base method
func (m *MockMinibrokerClient) Deprovision(arg0 string, arg1 bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deprovision", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deprovision indicates an expected call of Deprovision
func (mr *MockMinibrokerClientMockRecorder) Deprovision(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deprovision", reflect.TypeOf((*MockMinibrokerClient)(nil).Deprovision), arg0, arg1)
}

// GetBinding mocks base method
func (m *MockMinibrokerClient) GetBinding(arg0, arg1 string) (*v2.GetBindingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBinding", arg0, arg1)
	ret0, _ := ret[0].(*v2.GetBindingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBinding indicates an expected call of GetBinding
func (mr *MockMinibrokerClientMockRecorder) GetBinding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBinding", reflect.TypeOf((*MockMinibrokerClient)(nil).GetBinding), arg0, arg1)
}

// Init mocks base method
func (m *MockMinibrokerClient) Init(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockMinibrokerClientMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockMinibrokerClient)(nil).Init), arg0)
}

// LastBindingOperationState mocks base method
func (m *MockMinibrokerClient) LastBindingOperationState(arg0, arg1 string) (*v2.LastOperationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastBindingOperationState", arg0, arg1)
	ret0, _ := ret[0].(*v2.LastOperationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastBindingOperationState indicates an expected call of LastBindingOperationState
func (mr *MockMinibrokerClientMockRecorder) LastBindingOperationState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastBindingOperationState", reflect.TypeOf((*MockMinibrokerClient)(nil).LastBindingOperationState), arg0, arg1)
}

// LastOperationState mocks base method
func (m *MockMinibrokerClient) LastOperationState(arg0 string, arg1 *v2.OperationKey) (*v2.LastOperationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastOperationState", arg0, arg1)
	ret0, _ := ret[0].(*v2.LastOperationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastOperationState indicates an expected call of LastOperationState
func (mr *MockMinibrokerClientMockRecorder) LastOperationState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastOperationState", reflect.TypeOf((*MockMinibrokerClient)(nil).LastOperationState), arg0, arg1)
}

// ListServices mocks base method
func (m *MockMinibrokerClient) ListServices() ([]v2.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices")
	ret0, _ := ret[0].([]v2.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockMinibrokerClientMockRecorder) ListServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockMinibrokerClient)(nil).ListServices))
}

// Provision mocks base method
func (m *MockMinibrokerClient) Provision(arg0, arg1, arg2, arg3 string, arg4 bool, arg5 map[string]interface{}) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provision", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Provision indicates an expected call of Provision
func (mr *MockMinibrokerClientMockRecorder) Provision(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provision", reflect.TypeOf((*MockMinibrokerClient)(nil).Provision), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Unbind mocks base method
func (m *MockMinibrokerClient) Unbind(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unbind", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unbind indicates an expected call of Unbind
func (mr *MockMinibrokerClientMockRecorder) Unbind(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unbind", reflect.TypeOf((*MockMinibrokerClient)(nil).Unbind), arg0, arg1)
}
