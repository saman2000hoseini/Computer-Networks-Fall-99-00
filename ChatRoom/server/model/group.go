package model

type Group struct {
	id      string `gorm:"primary_key;gorm:type:varchar(250)"`
	name    string `gorm:"type:varchar(250);not null"`
	members []uint `gorm:"null"`
}
