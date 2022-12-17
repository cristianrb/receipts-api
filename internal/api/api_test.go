package api

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	mockstorage "receipts-api/mocks/storage"
	"receipts-api/pkg/types"
	"strings"
	"testing"
)

func TestReceipts_AddReceipt(t *testing.T) {
	receipt := types.ReceiptRequest{
		Items: []int{1, 2},
	}
	body, _ := json.Marshal(receipt)
	req, _ := http.NewRequest(http.MethodPost, "/receipts", strings.NewReader(string(body)))
	rr := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	storage := mockstorage.NewMockReceiptStorage(ctrl)
	itemStorage := mockstorage.NewMockItemStorage(ctrl)
	expectedReceipt := types.Receipt{
		Id: 1,
		Items: []*types.Item{
			{Id: 1, ProductName: "Product name 1"},
			{Id: 2, ProductName: "Product name 2"},
		},
	}
	storage.EXPECT().CreateReceipt(&receipt).Return(&expectedReceipt, nil)

	rh := New(":5000", storage, itemStorage)
	handler := http.HandlerFunc(rh.AddReceipt)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusAccepted, rr.Code)
}

func TestReceipts_GetReceiptById(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/receipts/1", nil)
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	rr := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	storage := mockstorage.NewMockReceiptStorage(ctrl)
	itemStorage := mockstorage.NewMockItemStorage(ctrl)
	expectedReceipt := types.Receipt{
		Id: 1,
		Items: []*types.Item{
			{Id: 1, ProductName: "Product name 1"},
			{Id: 2, ProductName: "Product name 2"},
		},
	}
	storage.EXPECT().GetReceiptById(1).Return(&expectedReceipt, nil)

	rh := New(":5000", storage, itemStorage)
	handler := http.HandlerFunc(rh.GetReceiptById)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestReceipts_GetAllReceipts(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/receipts", nil)
	rr := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	storage := mockstorage.NewMockReceiptStorage(ctrl)
	itemStorage := mockstorage.NewMockItemStorage(ctrl)
	expectedReceipts := types.Receipts{
		{
			Id: 1,
			Items: []*types.Item{
				{Id: 1, ProductName: "Product name 1"},
				{Id: 2, ProductName: "Product name 2"},
			},
		},
		{
			Id: 2,
			Items: []*types.Item{
				{Id: 3, ProductName: "Product name 3"},
			},
		},
	}
	storage.EXPECT().GetAllReceipts().Return(expectedReceipts, nil)

	rh := New(":5000", storage, itemStorage)
	handler := http.HandlerFunc(rh.GetAllReceipts)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
