package phone

import "database/sql"

type Model struct {
	DB *sql.DB
}

func (m Model) Insert(phone *Phone) error {
	query := `
		INSERT INTO phones (parent_id, phone_number)
		VALUES ($1, $2)
		RETURNING phone_id
	`
	args := []interface{}{phone.ParentID, phone.PhoneNumber}
	return m.DB.QueryRow(query, args...).Scan(&phone.PhoneID)
}
