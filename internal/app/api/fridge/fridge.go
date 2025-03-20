package fridge

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"warehouse/internal/models"
)

type Storage interface {
	FridgeContent(ctx context.Context) ([]models.ProductQuantity, error)
	AddProduct(ctx context.Context, product string, quantity int32) error
	RemoveProduct(ctx context.Context, product string, quantity int32) error
}
type API struct {
	log    *slog.Logger
	fridge Storage
}

func New(storage Storage, log *slog.Logger) *API {
	return &API{
		log:    log,
		fridge: storage,
	}
}

func (f *API) FridgeContent(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "applications/json" {
		f.log.Error("Content-Type not application/json")
		http.Error(w, "invalid Content-Type", http.StatusBadRequest)
	}

	products, err := f.fridge.FridgeContent(r.Context())
	if err != nil {
		f.log.Error("failed to get fridge content", "error", err)
		http.Error(w, "failed to get fridge content", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(products)
	if err != nil {
		f.log.Error("failed to decode request", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (f *API) AddProduct(w http.ResponseWriter, r *http.Request) {

	var productReq struct {
		Product  string `json:"product"`
		Quantity int32  `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&productReq); err != nil {
		f.log.Error("failed to decode request", "error", err)
		http.Error(w, "invalid request format", http.StatusBadRequest)
		return
	}

	if productReq.Product == "" || productReq.Quantity <= 0 {
		http.Error(w, "invalid product data", http.StatusBadRequest)
		return
	}

	if err := f.fridge.AddProduct(r.Context(), productReq.Product, productReq.Quantity); err != nil {
		f.log.Error("failed to add product", "product", productReq.Product, "quantity", productReq.Quantity, "error", err)
		http.Error(w, "error adding product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (f *API) RemoveProduct(w http.ResponseWriter, r *http.Request) {
	var productReq struct {
		Product  string `json:"product"`
		Quantity int32  `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&productReq); err != nil {
		f.log.Error("failed to decode request", "error", err)
		http.Error(w, "invalid request format", http.StatusBadRequest)
		return
	}

	if productReq.Product == "" || productReq.Quantity <= 0 {
		http.Error(w, "invalid product data", http.StatusBadRequest)
		return
	}

	if err := f.fridge.RemoveProduct(r.Context(), productReq.Product, productReq.Quantity); err != nil {
		f.log.Error("failed to remove product", "product", productReq.Product, "quantity", productReq.Quantity, "error", err)
		http.Error(w, "error removing product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
