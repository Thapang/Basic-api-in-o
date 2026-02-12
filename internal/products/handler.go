package products

import (
	"log"
	"net/http"

	"github.com/Thapang/ecom/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	//products := []string{"hello", "world"}

	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(products)

	err := h.service.ListProduct(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := struct {
		Products []string `json:"products"`
	}{}

	json.Write(w, http.StatusOK, products)

}
