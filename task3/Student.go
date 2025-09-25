package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint `gorm:"primary_key;auto_increment:true"`
	Name  string
	Age   int
	Grade string
}

func Run(db *gorm.DB) {

	student1 := Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}

	students := []Student{}

	db.Create(&student1)
	db.Find(&students, "age>?", 18)
	fmt.Print(students)
	db.Find(&students, "name=?", "张三").Update("grade", "四年级")
	db.Delete(&students, "age<?", 15)

}
