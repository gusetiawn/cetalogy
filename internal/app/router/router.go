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

	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.DeleteUserByID).Methods("DELETE")

	r.HandleFunc("/categories", gigCategoryHandler.CreateGigCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", gigCategoryHandler.GetGigCategoryByID).Methods("GET")
	r.HandleFunc("/categories", gigCategoryHandler.GetGigCategories).Methods("GET")
	r.HandleFunc("/categories/{id}", gigCategoryHandler.DeleteGigCategoryByID).Methods("DELETE")

	return r
}
