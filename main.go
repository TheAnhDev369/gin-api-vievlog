package main

import (	
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)
func main()  {
	router := gin.Default()
	// Định nghĩa route
	router.GET("/", func(ctx * gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// Chạy server trên cổng
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}