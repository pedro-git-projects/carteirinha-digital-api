package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/pedro-git-projects/carteirinha-api/src/app"
)

func main() {
	application, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	application.RegisterRoutes()
	if err := application.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
