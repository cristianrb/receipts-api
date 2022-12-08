package api

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	mockstorage "receipts-api/mocks/storage"
	"receipts-api/pkg/types"
	"strings"
	"testing"
)

func TestReceipts_AddReceipt(t *testing.T) {
	receipt := types.Receipt{
		Items: []*types.Item{
			{ProductName: "Product name 1"},
			{ProductName: "Product name 2"},
		},
	}
	body, _ := json.Marshal(receipt)
	req, _ := http.NewRequest(http.MethodPost, "/receipts", strings.NewReader(string(body)))
	rr := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	storage := mockstorage.NewMockStorage(ctrl)
	expectedReceipt := types.Receipt{
		Id: 1,
		Items: []*types.Item{
			{Id: 1, ProductName: "Product name 1"},
			{Id: 2, ProductName: "Product name 2"},
		},
	}
	storage.EXPECT().CreateReceipt(&receipt).Return(&expectedReceipt, nil)

	rh := New(":5000", storage)
	handler := http.HandlerFunc(rh.AddReceipt)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusAccepted, rr.Code)
}
