// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/docker/libmachete/machete/spi/instance (interfaces: Provisioner)

package instance

import (
	spi "github.com/docker/libmachete/machete/spi"
	instance "github.com/docker/libmachete/machete/spi/instance"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Provisioner interface
type MockProvisioner struct {
	ctrl     *gomock.Controller
	recorder *_MockProvisionerRecorder
}

// Recorder for MockProvisioner (not exported)
type _MockProvisionerRecorder struct {
	mock *MockProvisioner
}

func NewMockProvisioner(ctrl *gomock.Controller) *MockProvisioner {
	mock := &MockProvisioner{ctrl: ctrl}
	mock.recorder = &_MockProvisionerRecorder{mock}
	return mock
}

func (_m *MockProvisioner) EXPECT() *_MockProvisionerRecorder {
	return _m.recorder
}

func (_m *MockProvisioner) Destroy(_param0 instance.ID) *spi.Error {
	ret := _m.ctrl.Call(_m, "Destroy", _param0)
	ret0, _ := ret[0].(*spi.Error)
	return ret0
}

func (_mr *_MockProvisionerRecorder) Destroy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Destroy", arg0)
}

func (_m *MockProvisioner) ListGroup(_param0 instance.GroupID) ([]instance.ID, *spi.Error) {
	ret := _m.ctrl.Call(_m, "ListGroup", _param0)
	ret0, _ := ret[0].([]instance.ID)
	ret1, _ := ret[1].(*spi.Error)
	return ret0, ret1
}

func (_mr *_MockProvisionerRecorder) ListGroup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListGroup", arg0)
}

func (_m *MockProvisioner) Provision(_param0 string) (*instance.ID, *spi.Error) {
	ret := _m.ctrl.Call(_m, "Provision", _param0)
	ret0, _ := ret[0].(*instance.ID)
	ret1, _ := ret[1].(*spi.Error)
	return ret0, ret1
}

func (_mr *_MockProvisionerRecorder) Provision(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Provision", arg0)
}