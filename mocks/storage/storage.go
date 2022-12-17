// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	types "receipts-api/pkg/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReceiptStorage is a mock of ReceiptStorage interface.
type MockReceiptStorage struct {
	ctrl     *gomock.Controller
	recorder *MockReceiptStorageMockRecorder
}

// MockReceiptStorageMockRecorder is the mock recorder for MockReceiptStorage.
type MockReceiptStorageMockRecorder struct {
	mock *MockReceiptStorage
}

// NewMockReceiptStorage creates a new mock instance.
func NewMockReceiptStorage(ctrl *gomock.Controller) *MockReceiptStorage {
	mock := &MockReceiptStorage{ctrl: ctrl}
	mock.recorder = &MockReceiptStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceiptStorage) EXPECT() *MockReceiptStorageMockRecorder {
	return m.recorder
}

// CreateReceipt mocks base method.
func (m *MockReceiptStorage) CreateReceipt(receipt *types.ReceiptRequest) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReceipt", receipt)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReceipt indicates an expected call of CreateReceipt.
func (mr *MockReceiptStorageMockRecorder) CreateReceipt(receipt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReceipt", reflect.TypeOf((*MockReceiptStorage)(nil).CreateReceipt), receipt)
}

// DeleteReceiptById mocks base method.
func (m *MockReceiptStorage) DeleteReceiptById(receiptId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReceiptById", receiptId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReceiptById indicates an expected call of DeleteReceiptById.
func (mr *MockReceiptStorageMockRecorder) DeleteReceiptById(receiptId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReceiptById", reflect.TypeOf((*MockReceiptStorage)(nil).DeleteReceiptById), receiptId)
}

// GetAllReceipts mocks base method.
func (m *MockReceiptStorage) GetAllReceipts() (types.Receipts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllReceipts")
	ret0, _ := ret[0].(types.Receipts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllReceipts indicates an expected call of GetAllReceipts.
func (mr *MockReceiptStorageMockRecorder) GetAllReceipts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllReceipts", reflect.TypeOf((*MockReceiptStorage)(nil).GetAllReceipts))
}

// GetReceiptById mocks base method.
func (m *MockReceiptStorage) GetReceiptById(receiptId int) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceiptById", receiptId)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceiptById indicates an expected call of GetReceiptById.
func (mr *MockReceiptStorageMockRecorder) GetReceiptById(receiptId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiptById", reflect.TypeOf((*MockReceiptStorage)(nil).GetReceiptById), receiptId)
}

// MockItemStorage is a mock of ItemStorage interface.
type MockItemStorage struct {
	ctrl     *gomock.Controller
	recorder *MockItemStorageMockRecorder
}

// MockItemStorageMockRecorder is the mock recorder for MockItemStorage.
type MockItemStorageMockRecorder struct {
	mock *MockItemStorage
}

// NewMockItemStorage creates a new mock instance.
func NewMockItemStorage(ctrl *gomock.Controller) *MockItemStorage {
	mock := &MockItemStorage{ctrl: ctrl}
	mock.recorder = &MockItemStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemStorage) EXPECT() *MockItemStorageMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockItemStorage) CreateItem(item *types.Item) (*types.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", item)
	ret0, _ := ret[0].(*types.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockItemStorageMockRecorder) CreateItem(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockItemStorage)(nil).CreateItem), item)
}

// DeleteItemById mocks base method.
func (m *MockItemStorage) DeleteItemById(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItemById", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteItemById indicates an expected call of DeleteItemById.
func (mr *MockItemStorageMockRecorder) DeleteItemById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItemById", reflect.TypeOf((*MockItemStorage)(nil).DeleteItemById), id)
}

// GetItems mocks base method.
func (m *MockItemStorage) GetItems(ids []any) (types.Items, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems", ids)
	ret0, _ := ret[0].(types.Items)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockItemStorageMockRecorder) GetItems(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockItemStorage)(nil).GetItems), ids)
}

// UpdateItem mocks base method.
func (m *MockItemStorage) UpdateItem(item *types.Item) (*types.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", item)
	ret0, _ := ret[0].(*types.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockItemStorageMockRecorder) UpdateItem(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockItemStorage)(nil).UpdateItem), item)
}
