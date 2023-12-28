package student

import (
	"time"
)

type Student struct {
	ID               int64     `json:"student_id"`
	CreatedAt        time.Time `json:"created_at"`
	AcademicRegister string    `json:"academic_register"`
	Name             string    `json:"name"`
	Sex              string    `json:"sex"`
	Hash             string    `json:"password"`
	Role             string    `json:"role"`
	ParentID         int64     `json:"parent_id"`
}
