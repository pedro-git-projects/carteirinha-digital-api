package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pedro-git-projects/carteirinha-api/src/data/student"
	"github.com/pedro-git-projects/carteirinha-api/src/validator"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"golang.org/x/crypto/bcrypt"
)

func (app *App) QRCodeHandler(ctx *gin.Context) {
	token, ok := ctx.Get("bearer_token")
	if !ok {
		ctx.String(http.StatusUnauthorized, "Authentication token not found in the context")
		return
	}

	parsedToken, err := jwt.ParseWithClaims(token.(string), &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.config.jwtSecret), nil
	})

	if err != nil || !parsedToken.Valid {
		ctx.String(http.StatusUnauthorized, "Invalid token")
		return
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok {
		ctx.String(http.StatusInternalServerError, "Failed to extract claims from token")
		return
	}

	studentID := claims.UserID

	qrc, err := qrcode.New(token.(string))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate QR code with error: %v", err)
		return
	}

	filepath := fmt.Sprintf("../assets/qr-code-%d.jpg", studentID)
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
}

func (app *App) createStudentHandler(c *gin.Context) {
	payload := student.CreateDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	s := student.New(payload.Username, string(hash))

	v := validator.New()
	payload.Validate(v)
	if !v.Valid() {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "failed validation check"})
		return
	}

	err = app.models.Students.Insert(s)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
			c.JSON(http.StatusBadRequest, gin.H{"error": "credentials taken"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := app.GenerateJWT(s.ID(), s.Role())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]any{
		"student": s,
		"token":   token,
	}

	c.JSON(http.StatusCreated, response)
}

func (app *App) signinStudentHandler(c *gin.Context) {
	payload := student.LoginDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v := validator.New()
	payload.Validate(v)
	if !v.Valid() {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "failed validation check"})
		return
	}

	student, err := app.models.Students.Authenticate(payload.Username, payload.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := app.GenerateJWT(student.ID(), student.Role())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"token": token,
	}

	c.JSON(http.StatusOK, response)
}
