package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Account struct {
	ID      uint `gorm:"primary_key"`
	Balance float64
}

type Transaction struct {
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

func AccountRun(db *gorm.DB) {

	//db.AutoMigrate(&Account{}, &Transaction{})
	//db.Create(&Account{
	//	Balance: 1000,
	//})
	//
	//db.Create(&Account{
	//	Balance: 1000,
	//})

	db.Transaction(func(tx *gorm.DB) error {

		a := Account{}
		db.Find(&a, "id=?", 1)
		if tx.Error != nil {
			return tx.Error
		}

		b := Account{}
		db.Find(&b, "id=?", 2)
		if tx.Error != nil {
			return tx.Error
		}
		fmt.Print(b)

		if a.Balance > 100 {
			db.Model(&a).Where("Account", "id=?", a.ID).Update("balance", a.Balance-100)
			db.Model(&b).Update("balance", b.Balance+100)
			transaction := Transaction{
				FromAccountId: a.ID,
				ToAccountId:   b.ID,
				Amount:        100,
			}
			db.Create(&transaction)
			fmt.Print("transaction")
		} else {
			db.Rollback()
			fmt.Println("<UNK>")
		}

		// 返回 nil 提交事务
		// return nil
		return nil
	})

}
