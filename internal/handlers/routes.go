package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (r *Receipts) Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Post("/receipts", r.AddReceipt)
	return mux
}
