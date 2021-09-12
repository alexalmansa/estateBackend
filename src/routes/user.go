package routes

import (
	"database/sql"
	model2 "estateBackend/src/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func UsersLogin(c *gin.Context) {
	user := model2.User{}
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
	user := model2.User{}
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
	passwordChange := model2.PasswordChange{}
	err := c.ShouldBindJSON(&passwordChange)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	isValid, i, err := isUservalid(c)

	if isValid && err == nil {
		err = passwordChange.ChangePassword(conn, i)
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

func GetMe(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	isValid, i, err := isUservalid(c)
	if isValid && err == nil {
		user := model2.User{}
		err, user := user.GetMyUser(conn, i)
		if err != nil {
			fmt.Println("Error in user.Register()" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error with the token",
		})
	}
}

func isUservalid(c *gin.Context) (bool, int, error) {
	bearer := c.Request.Header.Get("Authorization")
	fmt.Printf(bearer)
	split := strings.Split(bearer, "Bearer ")
	token := split[1]
	user := model2.User{}
	isValid, userId := user.IsTokenValid(token)
	i, err := strconv.Atoi(userId)
	return isValid, i, err
}
