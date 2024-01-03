package attendance

type RegisterStudentEntryDTO struct {
	StudentToken string `json:"student_token" binding:"required"`
}
