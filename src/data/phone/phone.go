package phone

type Phone struct {
	PhoneID     int64  `json:"phone_id"`
	ParentID    int64  `json:"parent_id"`
	PhoneNumber string `json:"phone_number"`
}
