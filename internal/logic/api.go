package api

func CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user.Create(username, password)
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func UpdateUser(c *gin.Context) {
	id := c.PostForm("id")
	username := c.PostForm("username")
	password := c.PostForm("password")
	user.Update(id, username, password)
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func DeleteUser(c *gin.Context) {
	id := c.PostForm("id")
	user.Delete(id)
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
