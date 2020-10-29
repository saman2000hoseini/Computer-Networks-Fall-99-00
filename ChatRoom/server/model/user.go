package model

type User struct {
	id       uint64   `gorm:"primary_key"`
	Username string   `gorm:"unique_index;not null"`
	Password string   `gorm:"not null"`
	email    *string  `gorm:"null"`
	groups   []string `gorm:"null"`
	friends  []uint64 `gorm:"null"`
}

func (u *User) CheckPassword(p string) bool {
	return p == u.Password
}
