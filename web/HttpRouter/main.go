package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type GormUser struct {
	ID       uint   `json:"id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	db, err := gorm.Open("mysql", "root:239096@(124.221.153.82:3310)/rag?"+"charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	//根据结构体创建一张表可以用db.AutoMigrate(&gormUser{})
	// user := GormUser{
	// 	Phone:    "15395691522",
	// 	Name:     "haoyananLiar",
	// 	Password: "6666666",
	// }

	// db.Save(&user)
	queryUser := new(GormUser)
	db.Where("Phone=15395691522").Find(&queryUser)
	fmt.Println(queryUser)
}
