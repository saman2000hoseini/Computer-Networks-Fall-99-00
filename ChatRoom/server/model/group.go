package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"strconv"
)

type Group struct {
	Name    string         `gorm:"size:250;primaryKey"`
	Admin   uint64         `gorm:"not null"`
	Members pq.StringArray `gorm:"type:text[]"`
}

func NewGroup(name string, admin uint64) *Group {
	return &Group{
		Name:    name,
		Admin:   admin,
		Members: []string{strconv.FormatUint(admin, 10)},
	}
}

type GroupRepo interface {
	Find(name string) (Group, error)
	Save(group *Group) error
	Update(group Group) error
}

type SQLGroupRepo struct {
	DB *gorm.DB
}

func (r SQLGroupRepo) Find(name string) (Group, error) {
	var stored Group
	err := r.DB.Where(&Group{Name: name}).First(&stored).Error

	return stored, err
}

func (r SQLGroupRepo) Save(group *Group) error {
	return r.DB.Create(group).Error
}

func (r SQLGroupRepo) Update(group Group) error {
	return r.DB.Save(group).Error
}
