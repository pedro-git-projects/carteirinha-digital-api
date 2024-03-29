package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pedro-git-projects/carteirinha-api/src/data/attendance"
	"github.com/pedro-git-projects/carteirinha-api/src/data/parent"
	"github.com/pedro-git-projects/carteirinha-api/src/data/phone"
	"github.com/pedro-git-projects/carteirinha-api/src/data/staff"
	"github.com/pedro-git-projects/carteirinha-api/src/data/student"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"golang.org/x/crypto/bcrypt"
)

func (app *App) healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "online"})
}

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

	ctx.FileAttachment(filepath, fmt.Sprintf("qr-code-%d.jpg", studentID))
}

func (app *App) createParentHandler(c *gin.Context) {
	payload := parent.CreateDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	p := &parent.Parent{
		Name:   payload.Name,
		Email:  payload.Email,
		Hash:   string(hash),
		Phones: make([]phone.Phone, len(payload.Phones)),
	}

	for i, phoneDTO := range payload.Phones {
		p.Phones[i] = phone.Phone{
			PhoneNumber: phoneDTO.PhoneNumber,
		}
	}

	err = app.models.Parents.Insert(p)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "parents_email_key"` {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email already taken"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"parent": p,
	}

	c.JSON(http.StatusCreated, response)
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

	s := &student.Student{
		AcademicRegister: payload.AcademicRegister,
		Name:             payload.Name,
		Sex:              payload.Sex,
		Hash:             string(hash),
		ParentID:         payload.ParentID,
	}

	err = app.models.Students.Insert(s)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "students_academic_register_key"` {
			c.JSON(http.StatusBadRequest, gin.H{"error": "academic register already taken"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"student": s,
	}

	c.JSON(http.StatusCreated, response)
}

func (app *App) createParentStudentHandler(c *gin.Context) {

	type CreateParentStudentDTO struct {
		Parent  parent.CreateDTO               `json:"parent" binding:"required"`
		Student student.CreateParentStudentDTO `json:"student" binding:"required"`
	}
	payload := CreateParentStudentDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parentHash, err := bcrypt.GenerateFromPassword([]byte(payload.Parent.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash parent password"})
		return
	}

	parent := &parent.Parent{
		Name:  payload.Parent.Name,
		Email: payload.Parent.Email,
		Hash:  string(parentHash),
	}

	if len(payload.Parent.Phones) > 0 {
		for _, phoneDTO := range payload.Parent.Phones {
			parent.Phones = append(parent.Phones, phone.Phone{
				PhoneNumber: phoneDTO.PhoneNumber,
			})
		}
	}
	err = app.models.Parents.Insert(parent)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "parents_email_key"` {
			c.JSON(http.StatusBadRequest, gin.H{"error": "parent email already taken"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	studentHash, err := bcrypt.GenerateFromPassword([]byte(payload.Student.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash student password"})
		return
	}

	student := &student.Student{
		AcademicRegister: payload.Student.AcademicRegister,
		Name:             payload.Student.Name,
		Sex:              payload.Student.Sex,
		Hash:             string(studentHash),
		ParentID:         parent.ID,
	}

	err = app.models.Students.Insert(student)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "students_academic_register_key"` {
			c.JSON(http.StatusBadRequest, gin.H{"error": "student academic register already taken"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"parent":  parent,
		"student": student,
	}

	c.JSON(http.StatusCreated, response)
}

func (app *App) signinStudentHandler(c *gin.Context) {
	payload := student.LoginDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, err := app.models.Students.Authenticate(payload.AcademicRegister, payload.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := app.GenerateJWT(student.ID, student.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"token": token,
	}

	c.JSON(http.StatusOK, response)
}

func (app *App) signinParentHandler(c *gin.Context) {
	payload := parent.LoginDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parent, err := app.models.Parents.Authenticate(payload.ID, payload.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := app.GenerateJWT(parent.ID, parent.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"token": token,
	}

	c.JSON(http.StatusOK, response)
}

func (app *App) createStaffHandler(c *gin.Context) {
	payload := staff.CreateDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	staff := &staff.Staff{
		Chapa:        payload.Chapa,
		Name:         payload.Name,
		Hash:         string(hash),
		FunctionName: payload.FunctionName,
		Whatsapp:     payload.Whatsapp,
	}

	err = app.models.Staff.Insert(staff)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "staff_chapa_key"` {
			c.JSON(http.StatusBadRequest, gin.H{"error": "chapa already taken"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"staff": staff,
	}

	c.JSON(http.StatusCreated, response)
}

func (app *App) signinStaffHandler(c *gin.Context) {
	payload := staff.SignInDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staff, err := app.models.Staff.Authenticate(payload.Chapa, payload.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := app.GenerateJWT(staff.ID, staff.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"token": token,
	}

	c.JSON(http.StatusOK, response)
}

func (app *App) registerSchoolEntry(c *gin.Context) {
	payload := attendance.RegisterStudentEntryDTO{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staffTokenString := c.GetHeader("Authorization")
	if staffTokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Staff token is missing"})
		return
	}

	studentClaims, err := app.validateToken(payload.StudentToken, "student")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid student token"})
		return
	}

	staffClaims, err := app.validateToken(staffTokenString, "staff")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid staff token"})
		return
	}

	if staffClaims.Role != "staff" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient privileges"})
		return
	}

	entry := attendance.Attendance{
		StudentID: studentClaims.UserID,
		StaffID:   staffClaims.UserID,
		EntryTime: time.Now(),
	}

	err = app.models.Attendance.Insert(&entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register student entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student entry registered successfully"})
}
