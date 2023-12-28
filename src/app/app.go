package app

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pedro-git-projects/carteirinha-api/src/data"
)

type App struct {
	config Config
	router *gin.Engine
	models data.Models
}

func NewApp() (*App, error) {
	cfg, err := NewConfig()
	if err != nil {
		return nil, err
	}

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	app := &App{
		config: *cfg,
		router: r,
	}

	return app, nil
}

func (app *App) Run() error {
	db, err := openDB(app.config)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			db.Close()
		}
	}()

	app.models = data.NewModels(db)
	fmt.Printf("Conexão com o banco de dados estabelecida\n")

	app.router.Run()
	return nil
}
