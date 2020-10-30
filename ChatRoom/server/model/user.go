package model

type User struct {
	Username string   `gorm:"primaryKey;not null"`
	Password string   `gorm:"not null"`
	Email    string   `gorm:"null"`
	Groups   []string `gorm:"type:text[];null"`
	Friends  []string `gorm:"type:text[];null"`
}

func (u *User) CheckPassword(p string) bool {
	return p == u.Password
}
