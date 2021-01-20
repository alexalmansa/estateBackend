package routes

import (
	"database/sql"
	"estateBackend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenterCreate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)

	renter := model.Renter{}
	c.ShouldBindJSON(&renter)
	err := renter.Create(conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, renter)
}
func GetRenter(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)

	buildingId, _ := c.GetQuery("renter_id")
	buildings, err := model.GetRenters(conn, buildingId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"buildings": buildings})
}
