package router

import (
	"github.com/cetalogy/internal/app/handler"
	"github.com/cetalogy/internal/app/service"

	"github.com/gorilla/mux"
)

func SetupRouter(userService *service.UserService, gigCategoryService *service.GigCategoryService) *mux.Router {
	r := mux.NewRouter()

	userHandler := handler.NewUserHandler(userService)
	gigCategoryHandler := handler.NewGigCategoryHandler(gigCategoryService)

	r.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/api/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", userHandler.DeleteUserByID).Methods("DELETE")

	r.HandleFunc("/api/categories", gigCategoryHandler.CreateGigCategory).Methods("POST")
	r.HandleFunc("/api/categories/{id}", gigCategoryHandler.GetGigCategoryByID).Methods("GET")
	r.HandleFunc("/api/categories", gigCategoryHandler.GetGigCategories).Methods("GET")
	r.HandleFunc("/api/categories/{id}", gigCategoryHandler.DeleteGigCategoryByID).Methods("DELETE")

	return r
}
