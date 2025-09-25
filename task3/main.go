package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gujin_t31?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	db := InitDB()
	//Run(db)
	//AccountRun(db)
	//EmployeeRun(db)

	RunHook(db)
}
