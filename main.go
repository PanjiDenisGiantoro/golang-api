package main

import (
	"github.com/gin-gonic/gin"
	"webapi2/handler"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/ping", handler.HandleRoot)
	v1.GET("/books/:id", handler.Handlebooks)
	v1.GET("/query", handler.HandleQuery)
	v1.POST("/books", handler.PostBookHandler)
	router.Run(":8080")

}
