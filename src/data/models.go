package data

import (
	"database/sql"
	"errors"

	"github.com/pedro-git-projects/carteirinha-api/src/data/student"
)

var ErrRecordNotFound = errors.New("Register not found")

type Models struct {
	Students student.Model
}

func NewModels(db *sql.DB) Models {
	return Models{
		Students: student.Model{DB: db},
	}
}
