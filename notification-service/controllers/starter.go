package controllers

import "github.com/gin-gonic/gin"

func StarterPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the starter page",
	})
}
