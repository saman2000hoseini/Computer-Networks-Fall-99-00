package model

type Group interface {
	Create(id uint, gid, name string)
}

type GroupRepo struct {
	id      string `gorm:"primary_key;gorm:type:varchar(250)"`
	name    string `gorm:"type:varchar(250);not null"`
	members []uint `gorm:"null"`
}
