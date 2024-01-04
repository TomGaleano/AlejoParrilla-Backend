package models

type Menu struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	IsAvailable bool    `json:"is_available"`
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (Menu) TableName() string {
	return "app_menu"
}
