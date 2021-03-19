package routes

import (
	"database/sql"
	"estateBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func UsersLogin(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	err = user.IsAuthenticated(conn)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "There was an error authenticating.",
	})
}
func UsersRegister(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	err = user.Register(conn)
	if err != nil {
		fmt.Println("Error in user.Register()")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.GetAuthToken()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error getting token ": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
func UsersChangePassword(c *gin.Context) {
	user := model.PasswordChange{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	bearer := c.Request.Header.Get("Authorization")
	fmt.Printf(bearer)
	split := strings.Split(bearer, "Bearer ")
	token := split[1]
	isValid, userId := model.IsTokenValid(token)
	i, err := strconv.Atoi(userId)

	if isValid && err == nil {
		err = user.ChangePassword(conn, i)
		if err != nil {
			fmt.Println("Error in user.Register()" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "error with the token",
	})
}
