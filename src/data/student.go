package data

import "encoding/json"

type Roles string

const (
	RoleStaff   Roles = "staff"
	RoleStudent Roles = "student"
)

type Student struct {
	email string
	hash  string
}

func (s *Student) MarshalJSON() ([]byte, error) {
	data := map[string]string{
		"email":    s.email,
		"password": s.hash,
	}
	return json.Marshal(data)
}

func NewStudent(email, hash string) *Student {
	return &Student{email: email, hash: hash}
}
