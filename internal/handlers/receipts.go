package handlers

import (
	"net/http"
	"receipts-api/internal/services"
	"receipts-api/internal/utils"
	"receipts-api/pkg/models"
)

type Receipts struct {
	receiptsService services.ReceiptsService
}

// New creates an instance of Receipts
func New(receiptsService services.ReceiptsService) *Receipts {
	return &Receipts{
		receiptsService: receiptsService,
	}
}

// AddReceipt adds a receipt in the database
func (r *Receipts) AddReceipt(writer http.ResponseWriter, req *http.Request) {
	receipt := models.Receipt{}
	if err := utils.ReadJSON(writer, req, &receipt); err != nil {
		utils.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	insertedReceipt, err := r.receiptsService.CreateReceipt(receipt)
	if err != nil {
		utils.ErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(writer, http.StatusAccepted, insertedReceipt)
}
