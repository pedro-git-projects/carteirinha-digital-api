package staff

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Model struct {
	DB *sql.DB
}

func (m Model) Insert(s *Staff) error {
	query := `
		INSERT INTO staff (chapa, hash, role, name, function_name, whatsapp)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`
	args := []interface{}{s.Chapa, s.Hash, "staff", s.Name, s.FunctionName, s.Whatsapp}
	return m.DB.QueryRow(query, args...).Scan(&s.ID, &s.CreatedAt)
}

func (m Model) Authenticate(chapa, password string) (*Staff, error) {
	query := `
        SELECT id, created_at, chapa, name, hash, role, function_name, whatsapp
        FROM staff
        WHERE chapa = $1
    `

	staff := Staff{}
	err := m.DB.QueryRow(query, chapa).Scan(
		&staff.ID,
		&staff.CreatedAt,
		&staff.Chapa,
		&staff.Name,
		&staff.Hash,
		&staff.Role,
		&staff.FunctionName,
		&staff.Whatsapp,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("Staff not found")
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(staff.Hash), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid password")
	}

	return &staff, nil
}
