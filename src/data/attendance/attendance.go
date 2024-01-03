package attendance

import "time"

type Attendance struct {
	ID        int64     `json:"id"`
	StudentID int64     `json:"student_id"`
	CreatedAt time.Time `json:"created_at"`
	EntryTime time.Time `json:"entry_time"`
}
