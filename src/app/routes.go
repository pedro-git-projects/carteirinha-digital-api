package app

func (app *App) RegisterRoutes() {
	app.router.GET("/qr-code", app.QRCodeHandler)
}
