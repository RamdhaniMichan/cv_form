package entity

import "time"

type Model struct {
	ID        uint       `gorm:"primary_key;unique" sql:"type:serial;"`
	CreatedAt time.Time  `sql:"type:timestamp without time zone;"`
	UpdatedAt time.Time  `sql:"type:timestamp without time zone;"`
	DeletedAt *time.Time `sql:"index;type:timestamp without time zone;"`
}
