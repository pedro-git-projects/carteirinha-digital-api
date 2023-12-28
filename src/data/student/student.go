package student

import (
	"encoding/json"
	"time"
)

type Student struct {
	id        int64
	createdAt time.Time
	username  string
	hash      string
	role      string
}

func (s Student) Username() string {
	return s.username
}

func (s Student) Hash() string {
	return s.hash
}

func (s *Student) MarshalJSON() ([]byte, error) {
	data := map[string]string{
		"username": s.username,
		"password": s.hash,
		"role":     string(s.role),
	}
	return json.Marshal(data)
}

func (s Student) ID() int64 {
	return s.id
}

func (s Student) Role() string {
	return s.role
}

func New(username, hash string) *Student {
	return &Student{username: username, hash: hash, role: "student"}
}
