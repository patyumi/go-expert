package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	//CategoryID int
	Categories []Category `gorm:"many2many:products_categories;"`
	//SerialNumber SerialNumber
	gorm.Model
}

/*
type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}*/

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{} /*, &SerialNumber{}*/)

	// lock otimista versiona quando alguém tá fazendo uma alteração
	// mais utilizados quando não tem muita concorrência

	// lock pessimista, locka a tabela do banco de dados e ninguém consegue usar
	// mais utilizados quando grandes aplicações tem grandes acessos / concorrência

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c).Error
	if err != nil {
		panic(err)
	}

	c.Name = "Bivolt"
	tx.Debug().Save(&c)
	tx.Commit()
}
