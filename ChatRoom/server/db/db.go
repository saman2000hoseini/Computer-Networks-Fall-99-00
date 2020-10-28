package db

import (
	"errors"
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/ChatRoom/server/model"
	"os"

	"github.com/jinzhu/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./myDB.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

//create tables from models
func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{}, &model.Group{}).Error
	return err
}

//make database and create tables for the first time
func FirstSetup() (*gorm.DB, error) {
	db, err := NewDB()
	if err != nil {
		return nil, errors.New("error on creating db")
	}
	err = migrate(db)
	if err != nil {
		return nil, errors.New("error on creating tables")
	}
	return db, nil
}

func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../myDB_test.db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	if err :=
		os.Remove("./../myDB_test.db"); err != nil {
		return err
	}
	return nil
}
