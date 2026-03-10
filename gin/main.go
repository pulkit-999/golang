package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Define a simple GET endpoint
	router.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/ping/:id", func(c *gin.Context) {
		id := c.Param("id")
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"id":      id,
		})
	})

	router.POST("/me", func(c *gin.Context) {

		type user struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password"`
		}

		var x user

		if err := c.BindJSON(&x); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// Return JSON response
		//fmt.Printf("--- Incoming Request ---\nEmail: %s\nPassword: %s\n------------------------\n", x.Email, x.Password)
		c.JSON(http.StatusOK, gin.H{
			"email":    x.Email,
			"password": x.Password,
		})

	})

	router.PUT("/me", func(c *gin.Context) {

		type user struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password"`
		}

		var x user

		if err := c.BindJSON(&x); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// Return JSON response
		//fmt.Printf("--- Incoming Request ---\nEmail: %s\nPassword: %s\n------------------------\n", x.Email, x.Password)
		c.JSON(http.StatusOK, gin.H{
			"email":    x.Email,
			"password": x.Password,
		})

	})

	router.DELETE("/me/:key", func(c *gin.Context) {
		var id = c.Param("key")

		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"message": "",
		})
	})
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	router.Run()
}
