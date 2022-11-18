package services

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"receipts-api/mocks/repositories"
	"receipts-api/pkg/models"
	"testing"
)

func TestCreateReceipt(t *testing.T) {
	ctrl := gomock.NewController(t)
	receiptsRepositoryMock := mock_repositories.NewMockReceiptsRepository(ctrl)

	receipt := models.Receipt{
		Items: []*models.Item{
			{ProductName: "Product name 1"},
			{ProductName: "Product name 2"},
		},
	}
	expectedReceipt := models.Receipt{
		Id: 1,
		Items: []*models.Item{
			{Id: 1, ProductName: "Product name 1"},
			{Id: 2, ProductName: "Product name 2"},
		},
	}
	receiptsRepositoryMock.EXPECT().CreateReceipt(&receipt).Return(&expectedReceipt, nil)

	subject := New(receiptsRepositoryMock)
	actualReceipt, err := subject.CreateReceipt(receipt)

	assert.Equal(t, expectedReceipt, *actualReceipt)
	assert.NoError(t, err)
}

func TestCreateReceiptThrowsAnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	receiptsRepositoryMock := mock_repositories.NewMockReceiptsRepository(ctrl)

	receipt := models.Receipt{
		Items: []*models.Item{
			{ProductName: "Product name 1"},
			{ProductName: "Product name 2"},
		},
	}
	receiptsRepositoryMock.EXPECT().CreateReceipt(&receipt).Return(nil, errors.New("Error"))

	subject := New(receiptsRepositoryMock)
	actualReceipt, err := subject.CreateReceipt(receipt)

	assert.Error(t, err)
	assert.Nil(t, actualReceipt)
}
