package routes

import (
	"estateBackend/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func FlatCreate( c *gin.Context){
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	flat := model.Flat{}
	c.ShouldBindJSON(&flat)
	err := flat.Create(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, flat)
}
