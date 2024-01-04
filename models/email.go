package models

type Email struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (Email) TableName() string {
	return "app_contact"
}
