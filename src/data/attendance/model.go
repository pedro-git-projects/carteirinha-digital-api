package attendance

import (
	"database/sql"
)

type Model struct {
	DB *sql.DB
}

func (m Model) Insert(attendance *Attendance) error {
	query := `
        INSERT INTO attendance (student_id, staff_id, entry_time)
        VALUES ($1, $2, $3)
        RETURNING id, entry_time
    `
	args := []interface{}{attendance.StudentID, attendance.StaffID, attendance.EntryTime}
	err := m.DB.QueryRow(query, args...).Scan(&attendance.ID, &attendance.EntryTime)
	if err != nil {
		return err
	}

	return nil
}
