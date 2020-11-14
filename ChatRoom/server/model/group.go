package model

import "github.com/lib/pq"

type Group struct {
	Name    string         `gorm:"size:250;primaryKey"`
	Admin   string         `gorm:"size:250;null"`
	Members pq.StringArray `gorm:"type:text[]"`
}
