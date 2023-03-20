package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuhin37/truecaller-prefix/prefix"
)

func CheckPrefix(c *gin.Context) {
	checkString := c.Param("input")
	matchedPrefix := prefix.CheckPrefix(checkString)
	if matchedPrefix != "" {
		c.AsciiJSON(http.StatusOK, gin.H{
			"status": "successful",
			"prefix": matchedPrefix,
		})
	} else {
		c.AsciiJSON(http.StatusOK, gin.H{
			"status": "error",
			"prefix": "",
		})
	}
}
