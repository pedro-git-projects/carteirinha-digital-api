package student

import (
	"github.com/pedro-git-projects/carteirinha-api/src/validator"
)

type CreateDTO struct {
	AcademicRegister string `json:"academic_register" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Sex              string `json:"sex" binding:"required,oneof=Masculino Feminino"`
	Password         string `json:"password" binding:"required,min=6"`
	ParentID         int64  `json:"parent_id" binding:"required"`
}

type CreateParentStudentDTO struct {
	AcademicRegister string `json:"academic_register" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Sex              string `json:"sex" binding:"required,oneof=Masculino Feminino"`
	Password         string `json:"password" binding:"required,min=6"`
}

// func (dto CreateDTO) Validate(v *validator.Validator) {
// 	v.Check(dto.Username != "", "username", "is required")
// 	v.Check(dto.Password != "", "password", "is required")
// }

type LoginDTO struct {
	AcademicRegister string `json:"academic_register"`
	Password         string `json:"password"`
}

func (dto LoginDTO) Validate(v *validator.Validator) {
	v.Check(dto.AcademicRegister != "", "academic_register", "is required")
	v.Check(dto.Password != "", "password", "is required")
}
