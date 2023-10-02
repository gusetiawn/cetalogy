package app

import (
	"github.com/cetalogy/internal/app/repository"
	"github.com/cetalogy/internal/app/router"
	"github.com/cetalogy/internal/app/service"
	"github.com/cetalogy/internal/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router      *mux.Router
	UserService *service.UserService
}

func NewApp() *App {
	database.InitializeFirestore()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	r := router.SetupRouter(userService)

	return &App{
		Router:      r,
		UserService: userService,
	}
}

func (app *App) Start() {
	http.Handle("/", app.Router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
