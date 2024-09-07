package service

import (
	"Gim/internal/logic"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserList godoc
// @Summary      GetUserList
// @Tags         User
// @Success      200  {string}  "username, password"
// @Router       /user/getUserList [get]
func GetUserList(c *gin.Context) {
	users := logic.GetUserList()
	for _, user := range users {
		c.JSON(http.StatusOK, gin.H{"username": user.Username, "password": user.Password})
	}
}

// CreateUser godoc
// @Summary      CreateUser
// @Tags         User
// @Param        username query string true "username"
// @Param        password query string true "password"
// @Param        repassword query string true "repassword"
// @Success      200  {string}  "username, password"
// @Router       /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := logic.UserInfo{}
	user.Username = c.Query("username")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(-1, gin.H{"msg": "password different"})
	}
	user.Password = password
	logic.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{"msg": "Create user success!"})
}

// DeleteUser godoc
// @Summary      DeleteUser
// @Tags         User
// @Param        username query string true "id"
// @Success      200  {string}  "username, password"
// @Router       /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := logic.UserInfo{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.Id = uint64(id)
	logic.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{"msg": "Delete user success!"})
}
