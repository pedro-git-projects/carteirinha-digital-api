package student

import (
	"github.com/pedro-git-projects/carteirinha-api/src/validator"
)

type CreateDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dto CreateDTO) Validate(v *validator.Validator) {
	v.Check(dto.Username != "", "username", "is required")
	v.Check(dto.Password != "", "password", "is required")
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dto LoginDTO) Validate(v *validator.Validator) {
	v.Check(dto.Username != "", "username", "is required")
	v.Check(dto.Password != "", "password", "is required")
}
