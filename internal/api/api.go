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
	"strings"
)

type Server struct {
	listenAddr     string
	receiptStorage storage.ReceiptStorage
	itemStorage    storage.ItemStorage
}

// New creates an instance of Server
func New(listenAddr string, receiptStorage storage.ReceiptStorage, itemStorage storage.ItemStorage) *Server {
	return &Server{
		listenAddr:     listenAddr,
		receiptStorage: receiptStorage,
		itemStorage:    itemStorage,
	}
}

func (s *Server) Run() {
	mux := mux.NewRouter()
	mux.HandleFunc("/receipts", s.AddReceipt).Methods(http.MethodPost)
	mux.HandleFunc("/receipts/{id}", s.GetReceiptById).Methods(http.MethodGet)
	mux.HandleFunc("/receipts/{id}", s.DeleteReceiptById).Methods(http.MethodDelete)
	mux.HandleFunc("/receipts", s.GetAllReceipts).Methods(http.MethodGet)

	mux.HandleFunc("/items", s.AddItem).Methods(http.MethodPost)
	mux.HandleFunc("/items", s.GetItems).Methods(http.MethodGet)
	mux.HandleFunc("/items/{id}", s.UpdateItem).Methods(http.MethodPut)
	mux.HandleFunc("/items/{id}", s.DeleteItem).Methods(http.MethodDelete)

	logger.Info(fmt.Sprintf("Started server at port: %s", s.listenAddr))
	http.ListenAndServe(s.listenAddr, mux)
}

// AddReceipt adds a receipt in the database
func (s *Server) AddReceipt(writer http.ResponseWriter, req *http.Request) {
	receipt := &types.ReceiptRequest{}
	if err := utils.ReadJSON(writer, req, receipt); err != nil || receipt.Items == nil {
		logger.Error("Error when reading json at AddReceipt handler", err)
		if err == nil && receipt.Items == nil {
			utils.ErrorJSON(writer, errors.New("invalid json"), http.StatusBadRequest)
		}
		return
	}

	insertedReceipt, err := s.receiptStorage.CreateReceipt(receipt)
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

	receipt, err := s.receiptStorage.GetReceiptById(receiptId)
	if err != nil {
		logger.Error("Error when calling GetReceiptById", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, receipt)
}

func (s *Server) GetAllReceipts(w http.ResponseWriter, req *http.Request) {
	receipt, err := s.receiptStorage.GetAllReceipts()
	if err != nil {
		logger.Error("Error when calling GetAllReceipts", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, receipt)
}

func (s *Server) DeleteReceiptById(w http.ResponseWriter, req *http.Request) {
	receiptIdStr := mux.Vars(req)["id"]
	receiptId, err := strconv.Atoi(receiptIdStr)
	if err != nil {
		logger.Error("Error when converting receipt id to int", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = s.receiptStorage.DeleteReceiptById(receiptId)
	if err != nil {
		logger.Error("Error when calling GetReceiptById", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

// AddItem adds an item in the database
func (s *Server) AddItem(writer http.ResponseWriter, req *http.Request) {
	item := &types.Item{}
	if err := utils.ReadJSON(writer, req, item); err != nil || item.ProductName == "" {
		logger.Error("Error when reading json at AddItem handler", err)
		if err == nil && item.ProductName == "" {
			utils.ErrorJSON(writer, errors.New("invalid json"), http.StatusBadRequest)
		}
		return
	}

	insertedItem, err := s.itemStorage.CreateItem(item)
	if err != nil {
		logger.Error("Error when calling CreateItem", err)
		utils.ErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(writer, http.StatusAccepted, insertedItem)
}

func (s *Server) GetItems(writer http.ResponseWriter, req *http.Request) {
	idsReq := req.URL.Query().Get("ids")
	if idsReq == "" {
		s.GetAllItems(writer, req)
		return
	}
	idsStr := strings.Split(idsReq, ",")
	var ids []any
	for _, idStr := range idsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			logger.Error("Error when converting id to int", err)
			utils.ErrorJSON(writer, err, http.StatusInternalServerError)
			return
		}
		ids = append(ids, id)
	}

	items, err := s.itemStorage.GetItems(ids)
	if err != nil {
		logger.Error("Error when calling GetItems", err)
		utils.ErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(writer, http.StatusOK, items)
}

func (s *Server) GetAllItems(writer http.ResponseWriter, req *http.Request) {
	items, err := s.itemStorage.GetAllItems()
	if err != nil {
		logger.Error("Error when calling GetAllItems", err)
		utils.ErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(writer, http.StatusOK, items)
}

func (s *Server) UpdateItem(writer http.ResponseWriter, request *http.Request) {
	itemIdStr := mux.Vars(request)["id"]
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		logger.Error("Error when converting item id to int", err)
		utils.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}
	item := &types.Item{}
	if err := utils.ReadJSON(writer, request, item); err != nil || item.ProductName == "" {
		logger.Error("Error when reading json at UpdateItem handler", err)
		if err == nil && item.ProductName == "" {
			utils.ErrorJSON(writer, errors.New("invalid json"), http.StatusBadRequest)
		}
		return
	}
	item.Id = int64(itemId)

	updatedItem, err := s.itemStorage.UpdateItem(item)
	if err != nil {
		logger.Error("Error when calling UpdateItem", err)
		utils.ErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(writer, http.StatusOK, updatedItem)
}

func (s *Server) DeleteItem(writer http.ResponseWriter, request *http.Request) {
	itemIdStr := mux.Vars(request)["id"]
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		logger.Error("Error when converting item id to int", err)
		utils.ErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	err = s.itemStorage.DeleteItemById(int64(itemId))
	if err != nil {
		logger.Error("Error when calling UpdateItem", err)
		utils.ErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(writer, http.StatusOK, nil)
}
