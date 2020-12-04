package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint64   `gorm:"primaryKey"`
	Username string   `gorm:"unique;not null"`
	Password string   `gorm:"not null"`
	Email    string   `gorm:"null"`
	Friends  []string `gorm:"type:text[];null"`
}

func NewUser(username, password, email string) *User {
	return &User{
		Username: username,
		Password: password,
		Email:    email,
		Friends:  nil,
	}
}

func (u *User) CheckPassword(p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	return err == nil
}

type UserRepo interface {
	Find(username string) (User, error)
	Save(user *User) error
	Update(user User) error
}

type SQLUserRepo struct {
	DB *gorm.DB
}

func (r SQLUserRepo) Find(username string) (User, error) {
	var stored User
	err := r.DB.Where(&User{Username: username}).First(&stored).Error

	return stored, err
}

func (r SQLUserRepo) Save(user *User) error {
	return r.DB.Create(user).Error
}

func (r SQLUserRepo) Update(user User) error {
	return r.DB.Save(user).Error
}
