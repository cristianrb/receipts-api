package services

import (
	"receipts-api/internal/repositories"
	"receipts-api/pkg/models"
)

var _ ReceiptsService = &ReceiptsServiceImpl{}

type ReceiptsService interface {
	CreateReceipt(receipt models.Receipt) (*models.Receipt, error)
}

type ReceiptsServiceImpl struct {
	receiptsRepository repositories.ReceiptsRepository
}

// New creates an instance of ReceiptsServiceImpl
func New(receiptsRepository repositories.ReceiptsRepository) *ReceiptsServiceImpl {
	return &ReceiptsServiceImpl{
		receiptsRepository: receiptsRepository,
	}
}

// CreateReceipt creates a receipt in the database
func (s *ReceiptsServiceImpl) CreateReceipt(receipt models.Receipt) (*models.Receipt, error) {
	return s.receiptsRepository.CreateReceipt(&receipt)
}
