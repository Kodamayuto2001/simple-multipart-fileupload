package main

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserInput struct {
	File    *multipart.FileHeader `form:"file" binding:"required"`
	Keyword string                `form:"keyword" binding:"required,min=3"`
}

var validate *validator.Validate

func main() {
	r := gin.Default()

	// 1 << 20 は 1MBを意味します。
	r.MaxMultipartMemory = 1 << 20

	r.POST("/upload", func(c *gin.Context) {
		var input UserInput

		if err := c.ShouldBind(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Success"})
	})

	r.Run()
}
