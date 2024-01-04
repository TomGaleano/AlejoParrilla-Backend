package models

type Booking struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name" gorm:"type:varchar(40)"`
	Phone      string `json:"phone" gorm:"type:varchar(12)"`
	Date       string `json:"date" gorm:"type:date"`
	Time       string `json:"time" gorm:"type:time"`
	Seats      int    `json:"seats"`
	Restaurant string `json:"restaurant" gorm:"type:varchar(1000)"`
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (Booking) TableName() string {
	return "app_booking"
}
