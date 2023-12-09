package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"webapi2/book"
)

func HandleRoot(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ponging",
	})
}

func Handlebooks(c *gin.Context) {
	c.JSON(200, gin.H{
		"id": c.Param("id"),
	})
}
func HandleQuery(c *gin.Context) {
	c.JSON(200, gin.H{
		"query": c.Query("query"),
	})
}

func PostBookHandler(c *gin.Context) {
	var bookInput book.BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}
