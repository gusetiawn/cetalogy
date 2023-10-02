package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cetalogy/internal/app/model"
	"github.com/cetalogy/internal/app/service"
)

type GigCategoryHandler struct {
	gigCategoryService *service.GigCategoryService
}

func NewGigCategoryHandler(us *service.GigCategoryService) *GigCategoryHandler {
	return &GigCategoryHandler{gigCategoryService: us}
}

func (uh *GigCategoryHandler) CreateGigCategory(w http.ResponseWriter, r *http.Request) {
	var req model.GigServiceCategory
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := uh.gigCategoryService.CreateGigCategory(r.Context(), &req); err != nil {
		http.Error(w, "Error creating gig category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}

func (uh *GigCategoryHandler) GetGigCategoryByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID := params["id"]

	result, err := uh.gigCategoryService.GetGigCategoryByID(r.Context(), categoryID)
	if err != nil {
		http.Error(w, "Gig Category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (uh *GigCategoryHandler) GetGigCategories(w http.ResponseWriter, r *http.Request) {

	result, err := uh.gigCategoryService.GetGigCategories(r.Context())
	if err != nil {
		http.Error(w, "Gig Category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (uh *GigCategoryHandler) DeleteGigCategoryByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID := params["id"]

	err := uh.gigCategoryService.DeleteGigCategoryByID(r.Context(), categoryID)
	if err != nil {
		http.Error(w, "Gig Category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success delete categoryID: " + categoryID)
}
