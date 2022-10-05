// Code generated by MockGen. DO NOT EDIT.
// Source: product-api/repository (interfaces: GormRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockGormRepo is a mock of GormRepo interface.
type MockGormRepo struct {
	ctrl     *gomock.Controller
	recorder *MockGormRepoMockRecorder
}

// MockGormRepoMockRecorder is the mock recorder for MockGormRepo.
type MockGormRepoMockRecorder struct {
	mock *MockGormRepo
}

// NewMockGormRepo creates a new mock instance.
func NewMockGormRepo(ctrl *gomock.Controller) *MockGormRepo {
	mock := &MockGormRepo{ctrl: ctrl}
	mock.recorder = &MockGormRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGormRepo) EXPECT() *MockGormRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockGormRepo) Create(arg0 interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockGormRepoMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGormRepo)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockGormRepo) Delete(arg0 interface{}, arg1 ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockGormRepoMockRecorder) Delete(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGormRepo)(nil).Delete), varargs...)
}

// Find mocks base method.
func (m *MockGormRepo) Find(arg0 interface{}, arg1 ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockGormRepoMockRecorder) Find(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockGormRepo)(nil).Find), varargs...)
}

// First mocks base method.
func (m *MockGormRepo) First(arg0 interface{}, arg1 ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockGormRepoMockRecorder) First(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockGormRepo)(nil).First), varargs...)
}

// Model mocks base method.
func (m *MockGormRepo) Model(arg0 interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model", arg0)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockGormRepoMockRecorder) Model(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockGormRepo)(nil).Model), arg0)
}

// Updates mocks base method.
func (m *MockGormRepo) Updates(arg0 interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", arg0)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Updates indicates an expected call of Updates.
func (mr *MockGormRepoMockRecorder) Updates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*MockGormRepo)(nil).Updates), arg0)
}

// Where mocks base method.
func (m *MockGormRepo) Where(arg0 interface{}, arg1 ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Where", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Where indicates an expected call of Where.
func (mr *MockGormRepoMockRecorder) Where(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*MockGormRepo)(nil).Where), varargs...)
}
