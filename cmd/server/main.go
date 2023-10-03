package main

import (
	"github.com/cetalogy/internal/app"
	"github.com/cetalogy/internal/database"
)

func main() {
	database.InitializeFirestore()
	app := app.NewApp()
	app.Start()
}
