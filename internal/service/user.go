package service

import (
	"Gim/internal/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userInfo struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}

// CreateUser godoc
// @Summary      CreateUser
// @Tags         User
// @Param        user body userInfo true "username, password, repassword"
// @Success      201  {string}  "Create user success!"
// @Failure      400  {string}  "Invalid input"
// @Failure      400  {string}  "Passwords do not match"
// @Failure      500  {string}  "Internal server error"
// @Router       /user/createUser [post]
func CreateUser(c *gin.Context) {
	var newUser userInfo
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
		return
	}
	// check if the password and repassword are the same
	if newUser.Password != newUser.Repassword {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Passwords do not match"})
		return
	}
	// create the user
	var user logic.UserInfo
	user.Username = newUser.Username
	user.Password = newUser.Password
	if err := logic.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "Create user success!"})
}

// GetUserList godoc
// @Summary      GetUserList
// @Tags         User
// @Success      200  {string}  "Get user list"
// @Router       /user/getUserList [get]
func GetUserList(c *gin.Context) {
	users := logic.GetUserList()
	for _, user := range users {
		c.JSON(http.StatusOK, gin.H{"username": user.Username, "password": user.Password})
	}
}

// UpdateUser godoc
// @Summary      UpdateUser
// @Tags         User
// @Param        user body userInfo true "username, old password, new password"
// @Success      204  "Update user success!"
// @Router       /user/updateUser [put]
func UpdateUser(c *gin.Context) {
	var newUser userInfo
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
		return
	}
	// check if the user exists and the password is correct
	user, err := logic.GetUser(newUser.Username, newUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid username or password"})
		return
	}
	// update the password
	user.Password = newUser.Repassword
	if err := logic.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"msg": "Update user success!"})
}

// DeleteUser godoc
// @Summary      DeleteUser
// @Tags         User
// @Param        user body userInfo true "username, password, any"
// @Success      204  "Delete user success!"
// @Router       /user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	var checkUser userInfo
	if err := c.ShouldBindJSON(&checkUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
		return
	}
	// check if the user exists and the password is correct
	user, err := logic.GetUser(checkUser.Username, checkUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid username or password"})
		return
	}
	// delete the user
	if err := logic.DeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"msg": "Delete user success!"})
}
