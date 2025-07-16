package main

import (
	"GO_study/basic3"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Parent struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Child struct {
	Parent
	Age int
}

func InitDB(dst ...interface{}) *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(dst...)

	return db
}

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	//sqlx
	// db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic(err)
	// }

	basic3.Run_4(db)
	// basic3.Run_2(db)

	// lesson02.Run(db)
	// lesson03.Run(db)
	// lesson03_02.Run(db)
	// lesson03_03.Run(db)
	// lesson03_04.Run(db)
	// lesson04.Run(db)

	// InitDB(&Parent{}, &Child{})
}
