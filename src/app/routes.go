package app

import "github.com/gin-contrib/cors"

func (app *App) RegisterRoutes() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	app.router.Use(cors.New(config))

	authGroup := app.router.Group("/auth")
	authGroup.Use(app.BearerTokenMiddleware())
	authGroup.Use(app.TokenValidationMiddleware(app.config.jwtSecret))
	{
		authGroup.GET("/qr-code", app.QRCodeHandler)
	}

	app.router.POST("auth/students/signin", app.signinStudentHandler)
	app.router.POST("auth/parents/signin", app.signinParentHandler)
	app.router.POST("students", app.createStudentHandler)
	app.router.POST("parents", app.createParentHandler)
}
