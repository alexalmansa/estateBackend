package routes

import (
	"estateBackend/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func LeaseCreate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	lease := model.Lease{}
	c.ShouldBindJSON(&lease)
	err := lease.Create(&conn)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lease)
}

func GetLeases(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	flatId, _ := c.GetQuery("flat_id")
	renterid, _ := c.GetQuery("renter_id")

	leases, err := model.GetAllLeases(&conn, flatId, renterid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": leases})
}
