package service

import "github.com/gin-gonic/gin"

// GetIndex godoc
// @Tags         Index
// @Success      200  {string}  "Welcome!
// @Router       /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Welcome!",
	})
}
