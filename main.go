package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.Static("/assets", "./assets/")
	r.GET("qr-code", func(ctx *gin.Context) {
		a := Aluno{"pedro_meu_email@gmail.com", "falksjfdkaj12j31lj2"}

		jsn, err := a.MarshalJSON()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to marshal struct with error: %v", err)
			return
		}

		qrc, err := qrcode.New(string(jsn))
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to generate QR code with error: %v", err)
			return
		}

		filepath := "./assets/qr-code.jpg"
		w, err := standard.New(filepath)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to create writer with: %v", err)
			return
		}

		if err = qrc.Save(w); err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to write QR code to file with: %v", err)
			return
		}

		ctx.FileAttachment(filepath, "qrcode.jpg")
	})
	r.Run(":8080")
}
