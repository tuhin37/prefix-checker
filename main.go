package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuhin37/truecaller-prefix/controller"
)

func main() {
	r := gin.Default()
	r.GET("/check-prefix/:input", controller.CheckPrefix)

	r.GET("/health", func(c *gin.Context) {
		c.AsciiJSON(http.StatusOK, gin.H{
			"app":     "prefix-checker",
			"status":  "healthy",
			"version": "1.0.0",
		})
	})

	r.Run(":5000")
	// fmt.Println(prefix.CheckPrefix("humanknowledgeblongstotheworlddragonThunderdrag"))

}
