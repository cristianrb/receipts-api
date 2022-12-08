package api

import (
	"github.com/gorilla/mux"
	"net/http"
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

	http.ListenAndServe(s.listenAddr, mux)
}

// AddReceipt adds a receipt in the database
func (s *Server) AddReceipt(writer http.ResponseWriter, req *http.Request) {
	receipt := types.Receipt{}
	if err := utils.ReadJSON(writer, req, &receipt); err != nil {
		utils.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	insertedReceipt, err := s.mysqlStorage.CreateReceipt(&receipt)
	if err != nil {
		utils.ErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(writer, http.StatusAccepted, insertedReceipt)
}

func (s *Server) GetReceiptById(w http.ResponseWriter, req *http.Request) {
	receiptIdStr := mux.Vars(req)["id"]
	receiptId, err := strconv.Atoi(receiptIdStr)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	receipt, err := s.mysqlStorage.GetReceiptById(receiptId)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, receipt)
}

func (s *Server) GetAllReceipts(w http.ResponseWriter, req *http.Request) {
	receipt, err := s.mysqlStorage.GetAllReceipts()
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, receipt)
}
