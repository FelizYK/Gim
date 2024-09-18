package service

import (
	"Gim/internal/logic"
	"Gim/internal/utils"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type createUserInfo struct {
	Username   string
	Password   string
	RePassword string
}

// CreateUser godoc
//	@Summary	CreateUser
//	@Tags		User
//	@Param		user	body		createUserInfo	true	"username, password, repassword"
//	@Success	201		{string}	string			"Create user success!"
//	@Failure	400		{string}	string			"Invalid input"
//	@Failure	400		{string}	string			"Username already exists"
//	@Failure	400		{string}	string			"Passwords not same"
//  @Failure	500		{string}	string			"Random salt generation failed"
//	@Failure	500		{string}	string			"Internal server error"
//	@Router		/user/createUser [post]
func CreateUser(c *gin.Context) {
	var newUser createUserInfo
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
		return
	}
	// check if the username exists
	user, err := logic.GetUserByName(newUser.Username)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Username already exists"})
		return
	}
	// check if the password and repassword are the same
	if newUser.Password != newUser.RePassword {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Passwords not same"})
		return
	}
	// create
	user.Username = newUser.Username
	// user.Password = newUser.Password
	saltInt, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Random salt generation failed"})
		return
	}
	user.Salt = fmt.Sprintf("%x", saltInt)
	user.Password = utils.MakePassword(newUser.Password, user.Salt)
	if err := logic.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "Create user success!"})
}

// GetUserList godoc
//	@Summary	GetUserList
//	@Tags		User
//	@Success	200	{string}	string	"Get user list"
//	@Router		/user/getUserList [get]
func GetUserList(c *gin.Context) {
	users := logic.GetUserList()
	c.JSON(http.StatusOK, gin.H{"users": users})
}

type updateUserInfo struct {
	Username  string
	Password  string
	Parameter string
	Data      string
}

// UpdateUser godoc
//	@Summary	UpdateUser
//	@Tags		User
//	@Param		user	body	updateUserInfo	true	"username, password, password/telephone/email, data"
//	@Success	204		"Update user success!"
//	@Failure	400		{string}	string	"Invalid input"
//	@Failure	401		{string}	string	"Invalid username or password"
//	@Failure	400		{string}	string	"Invalid email"
//	@Failure	400		{string}	string	"Invalid parameter"
//	@Failure	500		{string}	string	"Internal server error"
//	@Router		/user/updateUser [put]
func UpdateUser(c *gin.Context) {
	// check login
	var newUser updateUserInfo
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
		return
	}
	user, err := logic.GetUserByName(newUser.Username)
	if err != nil || user.Password != newUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid username or password"})
		return
	}
	// check update
	switch newUser.Parameter {
	case "password":
		user.Password = newUser.Data
	case "telephone":
		user.Telephone = newUser.Data
	case "email":
		if !govalidator.IsEmail(newUser.Data) {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid email"})
			return
		}
		user.Email = newUser.Data
	default:
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid parameter"})
		return
	}
	// update
	if err := logic.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"msg": "Update user success!"})
}

type deleteUserInfo struct {
	Username string
	Password string
}

// DeleteUser godoc
//	@Summary	DeleteUser
//	@Tags		User
//	@Param		user	body	deleteUserInfo	true	"username, password"
//	@Success	204		"Delete user success!"
//	@Failure	400		{string}	string	"Invalid input"
//	@Failure	401		{string}	string	"Invalid username or password"
//	@Failure	500		{string}	string	"Internal server error"
//	@Router		/user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	// check if the user exists and the password is correct
	var newUser deleteUserInfo
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
		return
	}
	user, err := logic.GetUserByName(newUser.Username)
	if err != nil || user.Password != newUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid username or password"})
		return
	}
	// delete
	if err := logic.DeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"msg": "Delete user success!"})
}
