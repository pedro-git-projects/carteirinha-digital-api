package parent

import (
	"database/sql"
	"errors"

	"github.com/pedro-git-projects/carteirinha-api/src/data/phone"
	"golang.org/x/crypto/bcrypt"
)

type Model struct {
	DB         *sql.DB
	PhoneModel *phone.Model
}

func (m Model) Insert(parent *Parent) error {
	query := `
		INSERT INTO parents (created_at, name, email, hash, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`
	args := []interface{}{parent.CreatedAt, parent.Name, parent.Email, parent.Hash, "parent"}
	err := m.DB.QueryRow(query, args...).Scan(&parent.ID, &parent.CreatedAt)
	if err != nil {
		return err
	}

	for _, phone := range parent.Phones {
		phone.ParentID = parent.ID
		err := m.PhoneModel.Insert(&phone)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m Model) Authenticate(email, password string) (*Parent, error) {
	query := `
        SELECT id, created_at, name, email, hash, role
        FROM parents
        WHERE id = $1
    `

	parent := Parent{}
	err := m.DB.QueryRow(query, email).Scan(
		&parent.ID,
		&parent.CreatedAt,
		&parent.Name,
		&parent.Email,
		&parent.Hash,
		&parent.Role,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("Parent not found")
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(parent.Hash), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid password")
	}

	return &parent, nil
}
