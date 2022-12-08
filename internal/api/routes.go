package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) Routes() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/receipts", s.AddReceipt)
	mux.HandleFunc("/receipts/{id}", s.GetReceiptById)
	return mux
}
