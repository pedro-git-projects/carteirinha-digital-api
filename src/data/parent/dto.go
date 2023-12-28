package parent

import "github.com/pedro-git-projects/carteirinha-api/src/validator"

type PhoneCreateDTO struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type CreateDTO struct {
	Name     string           `json:"name" binding:"required"`
	Email    string           `json:"email" binding:"required,email"`
	Password string           `json:"password" binding:"required,min=6"`
	Phones   []PhoneCreateDTO `json:"phones"`
}

type LoginDTO struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func (dto LoginDTO) Validate(v *validator.Validator) {
	v.Check(dto.ID != "", "ID", "is required")
	v.Check(dto.Password != "", "password", "is required")
}
