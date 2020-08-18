package routes

import (
	"estateBackend/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func BuildingCreate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	building := model.Building{}
	c.ShouldBindJSON(&building)
	err := building.Create(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, building)
}

func GetBuildings(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	buildingId, _ := c.GetQuery("building_id")
	buildings, err := model.GetBuildings(&conn, buildingId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"buildings": buildings})
}
