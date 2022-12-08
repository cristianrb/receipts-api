package api

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"receipts-api/internal/logger"
	"receipts-api/internal/storage"
	"receipts-api/internal/utils"
	"receipts-api/pkg/types"
	"strconv"
)

type Server struct {
	listenAddr   string
	mysqlStorage storage.Storage
}

// New creates an instance of Server
func New(listenAddr string, mysqlStorage storage.Storage) *Server {
	return &Server{
		listenAddr:   listenAddr,
		mysqlStorage: mysqlStorage,
	}
}

func (s *Server) Run() {
	mux := mux.NewRouter()
	mux.HandleFunc("/receipts", s.AddReceipt).Methods(http.MethodPost)
	mux.HandleFunc("/receipts/{id}", s.GetReceiptById).Methods(http.MethodGet)
	mux.HandleFunc("/receipts", s.GetAllReceipts).Methods(http.MethodGet)

	logger.Info(fmt.Sprintf("Started server at port: %s", s.listenAddr))
	http.ListenAndServe(s.listenAddr, mux)
}

// AddReceipt adds a receipt in the database
func (s *Server) AddReceipt(writer http.ResponseWriter, req *http.Request) {
	receipt := &types.Receipt{}
	if err := utils.ReadJSON(writer, req, receipt); err != nil || receipt.Items == nil {
		logger.Error("Error when reading json at AddReceipt handler", err)
		if err == nil && receipt.Items == nil {
			utils.ErrorJSON(writer, errors.New("invalid json"), http.StatusBadRequest)
		}
		return
	}

	insertedReceipt, err := s.mysqlStorage.CreateReceipt(receipt)
	if err != nil {
		logger.Error("Error when calling CreateReceipt", err)
		utils.ErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(writer, http.StatusAccepted, insertedReceipt)
}

func (s *Server) GetReceiptById(w http.ResponseWriter, req *http.Request) {
	receiptIdStr := mux.Vars(req)["id"]
	receiptId, err := strconv.Atoi(receiptIdStr)
	if err != nil {
		logger.Error("Error when converting receipt id to int", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	receipt, err := s.mysqlStorage.GetReceiptById(receiptId)
	if err != nil {
		logger.Error("Error when calling GetReceiptById", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, receipt)
}

func (s *Server) GetAllReceipts(w http.ResponseWriter, req *http.Request) {
	receipt, err := s.mysqlStorage.GetAllReceipts()
	if err != nil {
		logger.Error("Error when calling GetAllReceipts", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, receipt)
}
