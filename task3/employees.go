package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Employee struct {
	Id         int    `db:"id" json:"id"`
	Name       string ` db:"name" json:"name"`
	Department string `db:"department" json:"department"`
	Salary     int    `db:"salary"  json:"salary"`
}

type Book struct {
	Id     int    `db:"id" json:"id"`
	Title  string ` db:"title" json:"title"`
	Author string `db:"author" json:"author"`
	Price  int    `db:"price"  json:"price"`
}

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gujin_t31"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("连接数据库失败")
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	fmt.Println("连接数据库成功")
	return nil
}

func EmployeeRun(gromDB *gorm.DB) {
	initDB()
	gromDB.AutoMigrate(&Book{})
	gromDB.Create(&Book{
		Title:  "book1",
		Author: "yang",
		Price:  1000,
	})

	gromDB.Create(&Book{
		Title:  "book2",
		Author: "yang2",
		Price:  1,
	})

	gromDB.Create(&Book{
		Title:  "book3",
		Author: "yang3",
		Price:  300,
	})

	var emps2 []Employee = make([]Employee, 100)
	err := db.Select(&emps2, "select * from employees where department=? ", "技术部")
	if err != nil {
		fmt.Println("查询失败")
		panic(err)
		return
	}
	fmt.Println(emps2)

	var topSalary Employee
	err = db.Get(&topSalary, "select * from employees order by salary desc limit 1")
	if err != nil {
		fmt.Println("查询失败")
		panic(err)
		return
	}
	fmt.Println(topSalary)

	var books []Book
	err = db.Select(&books, "select * from books where price > ?", 50)
	if err != nil {
		fmt.Println("<UNK>")
		panic(err)
		return
	}
	fmt.Println(books)

}
