package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"webapi2/book"
	"webapi2/handler"
)

func main() {
	dsn := "root:W@rung01@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&book.Book{})

	//create

	//book := book.Book{}
	//book.Title = "Golang 2"
	//book.Price = 20000
	//book.Discount = 30
	//book.Rating = 9
	//book.Description = "Golang is the best wkwkw"
	//
	//err = db.Create(&book).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//read
	//book := book.Book{}

	//find by id
	//var book book.Book
	//err = db.Debug().First(&book).Error

	//get
	var books []book.Book
	err = db.Debug().Find(&books).Error
	if err != nil {
		fmt.Println(err)
	}

	for _, b := range books {
		fmt.Println("title", b.Title)
	}

	//update
	//var book book.Book
	//
	//err = db.Debug().Where("id = ?", 1).First(&book).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//book.Title = "Golang 2"
	//book.Price = 20000
	//book.Discount = 30
	//book.Rating = 9
	//book.Description = "Golang is the best wkwkw"
	//err = db.Save(&book).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//delete
	//var book book.Book
	//
	//err = db.Debug().Where("id = ?", 1).Delete(&book).Error
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println("title", book.Title)
	//fmt.Println("price %v", book)

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/ping", handler.HandleRoot)
	v1.GET("/books/:id", handler.Handlebooks)
	v1.GET("/query", handler.HandleQuery)
	v1.POST("/books", handler.PostBookHandler)
	router.Run(":8080")

}
