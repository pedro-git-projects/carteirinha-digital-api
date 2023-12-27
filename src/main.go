package main

import (
	"log"

	"github.com/pedro-git-projects/carteirinha-api/src/app"
)

func main() {
	application, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	application.RegisterRoutes()
	application.Run()
}
