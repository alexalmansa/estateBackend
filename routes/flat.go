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
	flat := model.Flat{}

	flats, err := flat.GetBuildingItems(conn, buildingId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, flats)
}

func DeleteFlat(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	flatId, _ := c.GetQuery("flat_id")
	flat := model.Flat{}

	err := flat.DeleteFlat(conn, flatId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"flat " + flatId: "Deleted correctly"})
}

func UpdateFlat(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)

	flat := model.Flat{}
	c.ShouldBindJSON(&flat)
	err := flat.UpdateFlat(conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flat)
}
