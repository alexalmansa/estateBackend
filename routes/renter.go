package routes

import (
	"estateBackend/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func RenterCreate( c *gin.Context){
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	renter := model.Renter{}
	c.ShouldBindJSON(&renter)
	err := renter.Create(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, renter)
}
