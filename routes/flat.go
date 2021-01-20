package routes

import (
	"database/sql"
	"estateBackend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FlatCreate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)

	flat := model.Flat{}
	c.ShouldBindJSON(&flat)
	err := flat.Create(conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flat)
}

func FlatFromBuilding(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	buildingId, _ := c.GetQuery("building_id")

	flats, err := model.GetBuildingItems(conn, buildingId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"flats": flats})
}

func DeleteFlat(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	flatId, _ := c.GetQuery("flat_id")

	err := model.DeleteFlat(conn, flatId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"flat " + flatId: "Deleted correctly"})
}
