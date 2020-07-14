package routes

import (
	"estateBackend/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func LeaseCreate( c *gin.Context){
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	lease := model.Lease{}
	c.ShouldBindJSON(&lease)
	err := lease.Create(&conn)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, lease)
}