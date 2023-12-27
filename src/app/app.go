package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	config Config
	router *gin.Engine
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

func (app *App) Run() {
	app.router.Run()
}
