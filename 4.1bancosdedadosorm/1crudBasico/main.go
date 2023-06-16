package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm usa muitas tags

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	gorm.Model // soft delete,c ria campos de created, updated e deleted e trata isso
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create
	/*db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.00,
	})*/

	// create batch
	/*products := []Product{
		{Name: "Notebook", Price: 1000.00},
		{Name: "Mouse", Price: 50.00},
		{Name: "Keyboard", Price: 100.00},
	}
	db.Create(products)*/

	// buscar infos usando gorm
	//var product Product
	//db.First(&product, 1)
	//fmt.Println(product)

	//db.First(&product, "name = ?", "Mouse")
	//fmt.Println(product)

	// buscar tudo
	/*var products []Product
	db.Find(&products)

	for _, prod := range products {
		fmt.Println(prod)
	}*/

	// busca com limit
	/*
		var products []Product
		db.Limit(2).Offset(2).Find(&products)
		for _, prod := range products {
			fmt.Println(prod)
		}

		// trabalhando com where
		db.Where("price > ?", 100).Find(&products)
		for _, prod := range products {
			fmt.Println(prod)
		}
	*/

	/*
		var p Product
		db.First(&p, 1)
		p.Name = "New Mouse"
		db.Save(&p)
	*/

	/*
		var p2 Product
		db.First(&p2, 1)
		fmt.Println(p2.Name)

		db.Delete(&p2)
	*/
}
