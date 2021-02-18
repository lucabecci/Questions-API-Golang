package database

import (
	"fmt"
	"log"

	"github.com/lucabecci/questions-golang-API/pkg/question"
	"github.com/lucabecci/questions-golang-API/pkg/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Data struct {
	Db *gorm.DB
}

var database *gorm.DB

func GetInstance(uri string) (Data, error) {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), //logger blocked
	})

	if err != nil {
		log.Fatal("Error to connect your db")
	}

	data := Data{
		Db: db,
	}
	err = data.Migrations()
	if err != nil {
		log.Fatal("Error to create your migrations")
	}
	fmt.Println("DB is connected")
	database = data.Db
	return data, nil
}

func (d *Data) Migrations() error {
	userTable := d.Db.Migrator().HasTable(&user.User{})
	questionTable := d.Db.Migrator().HasTable(&question.Question{})
	if userTable == false && questionTable == false {
		err := d.Db.AutoMigrate(&user.User{})
		if err != nil {
			return err
		}
		err = d.Db.AutoMigrate(&question.Question{})
		if err != nil {
			return err
		}
		fmt.Println("Models Created")
		return nil
	}
	return nil
}

func (d *Data) Close() error {
	d.Close()
	return nil
}

func Database() *gorm.DB {
	return database
}
