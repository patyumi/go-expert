package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	/*
		// create category
		category := Category{Name: "Eletronicos"}
		db.Create(&category)

		category2 := Category{Name: "Cozinha"}
		db.Create(&category2)

		// create product
		db.Create(&Product{
			Name:       "Microondas",
			Price:      1000.00,
			Categories: []Category{category, category2},
			//CategoryID: category.ID,
		})
	*/

	/*
		var categories []Category
		// many 2 many
		err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
		if err != nil {
			panic(err)
		}

		for _, cat := range categories {
			fmt.Println(cat.Name, ":")
			for _, prod := range cat.Products {
				fmt.Println("- ", prod.Name)
			}
		}
	*/

	/*
		var categories []Category
		// has many
		err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
		if err != nil {
			panic(err)
		}

		for _, cat := range categories {
			fmt.Println(cat.Name, ":")
			for _, prod := range cat.Products {
				fmt.Println("- ", prod.Name, category.Name)
			}
		}
	*/

	/*
		// create serial number
		db.Create(&SerialNumber{
			Number:    "123456",
			ProductID: 1,
		})

		var categories []Category
		// has many pegadinha
		// tem que vincular a informação a quem ele tem relação
		err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
		if err != nil {
			panic(err)
		}

		for _, cat := range categories {
			fmt.Println(cat.Name, ":")
			for _, prod := range cat.Products {
				fmt.Println("- ", prod.Name, "SN: ", prod.SerialNumber.Number)
			}
		}

		/*
			var products []Product
			// has one
			db.Preload("Category").Preload("SerialNumber").Find(&products)

			for _, product := range products {
				fmt.Println(product)
			}
	*/
}
