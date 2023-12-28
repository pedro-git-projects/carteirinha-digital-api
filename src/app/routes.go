package app

func (app *App) RegisterRoutes() {
	app.router.GET("/qr-code", app.QRCodeHandler)
	app.router.POST("students", app.createStudentHandler)
	app.router.POST("auth/students/signin", app.signinStudentHandler)
}
