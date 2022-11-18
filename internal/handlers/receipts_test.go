package handlers

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"receipts-api/mocks/services"
	"receipts-api/pkg/models"
	"strings"
	"testing"
)

func TestReceipts_AddReceipt(t *testing.T) {
	receipt := models.Receipt{
		Items: []*models.Item{
			{ProductName: "Product name 1"},
			{ProductName: "Product name 2"},
		},
	}
	body, _ := json.Marshal(receipt)
	req, _ := http.NewRequest(http.MethodPost, "/receipts", strings.NewReader(string(body)))
	rr := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	receiptsService := mock_services.NewMockReceiptsService(ctrl)
	expectedReceipt := models.Receipt{
		Id: 1,
		Items: []*models.Item{
			{Id: 1, ProductName: "Product name 1"},
			{Id: 2, ProductName: "Product name 2"},
		},
	}
	receiptsService.EXPECT().CreateReceipt(receipt).Return(&expectedReceipt, nil)

	rh := New(receiptsService)
	handler := http.HandlerFunc(rh.AddReceipt)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusAccepted, rr.Code)
}
