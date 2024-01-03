package attendance

import (
	"database/sql"
	"time"
)

type Model struct {
	DB *sql.DB
}

func (m Model) Insert(studentID int) (int, time.Time, error) {
	query := `
	INSERT INTO attendances (student_id, entry_time)
	VALUES ($1, $2)
	RETURNING id, entry_time
`
	args := []interface{}{studentID, time.Now()}
	var id int
	var entryTime time.Time
	err := m.DB.QueryRow(query, args...).Scan(&id, &entryTime)
	if err != nil {
		return 0, time.Time{}, err
	}

	return id, entryTime, nil
}
