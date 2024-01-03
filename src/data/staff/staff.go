package staff

import "time"

type Staff struct {
	ID           int64     `json:"student_id"`
	CreatedAt    time.Time `json:"created_at"`
	Chapa        string    `json:"chapa"`
	Name         string    `json:"name"`
	Hash         string    `json:"password"`
	Role         string    `json:"role"`
	FunctionName string    `json:"function_name"`
	Whatsapp     string    `json:"whatsapp"`
}
