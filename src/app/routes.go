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
		authGroup.POST("/record-attendance", app.registerSchoolEntry)
	}

	app.router.POST("auth/students/signin", app.signinStudentHandler)
	app.router.POST("auth/parents/signin", app.signinParentHandler)
	app.router.POST("auth/staff/signin", app.signinStaffHandler)

	app.router.POST("parent-student", app.createParentStudentHandler)
	app.router.POST("students", app.createStudentHandler)
	app.router.POST("parents", app.createParentHandler)
	app.router.POST("staff", app.createStaffHandler)

	app.router.GET("/", app.healthCheckHandler)
}
