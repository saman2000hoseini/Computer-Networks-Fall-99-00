package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username string   `gorm:"primaryKey;not null"`
	Password string   `gorm:"not null"`
	Email    string   `gorm:"null"`
	Groups   []string `gorm:"type:text[];null"`
	Friends  []string `gorm:"type:text[];null"`
}

func (u *User) CheckPassword(p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	return err == nil
}
