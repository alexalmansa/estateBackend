package routes

import (
	"database/sql"
	model2 "estateBackend/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenterCreate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)

	renter := model2.Renter{}
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

	renterId, _ := c.GetQuery("renter_id")
	renter := model2.Renter{}

	renters, err := renter.GetRenters(conn, renterId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, renters)
}
func DeleteRenter(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	renterId, _ := c.GetQuery("renter_id")
	renter := model2.Renter{}

	err := renter.DeleteRenter(conn, renterId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Renter " + renterId: "Deleted correctly"})
}
func UpdateRenter(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)

	renter := model2.Renter{}
	c.ShouldBindJSON(&renter)
	err := renter.UpdateRenter(conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, renter)
}
