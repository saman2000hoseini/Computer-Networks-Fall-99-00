package model

type Group struct {
	Name    string   `gorm:"size:250;primaryKey"`
	Admin   string   `gorm:"size:250;not_null"`
	Members []string `gorm:"type:text[];null"`
}
