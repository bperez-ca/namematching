package http

import (
	"NameMatching/internal/app"
	"encoding/json"
	"net/http"
)

// HTTPAdapter implements the HTTPHandler interface
type HTTPAdapter struct {
	customerValidationService *app.CustomerValidationService
}

func NewHTTPAdapter(service *app.CustomerValidationService) *HTTPAdapter {
	return &HTTPAdapter{customerValidationService: service}
}

// NameMatchHandler handles name matching API requests
func (h *HTTPAdapter) NameMatchHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name1 string `json:"name1"`
		Name2 string `json:"name2"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)

	_, score := h.customerValidationService.ValidateCustomer(req.Name1, req.Name2, "", "", 0.8)
	err := json.NewEncoder(w).Encode(map[string]float64{"score": score})
	if err != nil {
		return
	}
}

// EmailMatchHandler handles email matching API requests
func (h *HTTPAdapter) EmailMatchHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email1 string `json:"email1"`
		Email2 string `json:"email2"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)

	_, score := h.customerValidationService.ValidateCustomer("", "", req.Email1, req.Email2, 0.8)
	err := json.NewEncoder(w).Encode(map[string]float64{"score": score})
	if err != nil {
		return
	}
}
