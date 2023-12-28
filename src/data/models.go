package data

import (
	"database/sql"
	"errors"

	"github.com/pedro-git-projects/carteirinha-api/src/data/parent"
	"github.com/pedro-git-projects/carteirinha-api/src/data/phone"
	"github.com/pedro-git-projects/carteirinha-api/src/data/student"
)

var ErrRecordNotFound = errors.New("Register not found")

type Models struct {
	Students student.Model
	Parents  parent.Model
}

func NewModels(db *sql.DB) Models {
	pm := phone.Model{DB: db}
	return Models{
		Students: student.Model{DB: db},
		Parents: parent.Model{
			DB:         db,
			PhoneModel: &pm,
		},
	}
}
