package parent

import (
	"time"

	"github.com/pedro-git-projects/carteirinha-api/src/data/phone"
)

type Parent struct {
	ID        int64         `json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Hash      string        `json:"hash"`
	Role      string        `json:"role"`
	Phones    []phone.Phone `json:"phones"`
}
