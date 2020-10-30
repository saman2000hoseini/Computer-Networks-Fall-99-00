package model

type Group struct {
	ID      string   `gorm:"size:250;primaryKey"`
	Name    string   `gorm:"size:250;not_null"`
	Members []string `gorm:"type:text[];null"`
}
