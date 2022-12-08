// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	types "receipts-api/pkg/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// CreateReceipt mocks base method.
func (m *MockStorage) CreateReceipt(receipt *types.Receipt) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReceipt", receipt)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReceipt indicates an expected call of CreateReceipt.
func (mr *MockStorageMockRecorder) CreateReceipt(receipt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReceipt", reflect.TypeOf((*MockStorage)(nil).CreateReceipt), receipt)
}

// GetReceiptById mocks base method.
func (m *MockStorage) GetReceiptById(receiptId int) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceiptById", receiptId)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceiptById indicates an expected call of GetReceiptById.
func (mr *MockStorageMockRecorder) GetReceiptById(receiptId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiptById", reflect.TypeOf((*MockStorage)(nil).GetReceiptById), receiptId)
}