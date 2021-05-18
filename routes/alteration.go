package routes

import (
	"database/sql"
	"estateBackend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AlterationCreate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)

	alteration := model.Alteration{}
	c.ShouldBindJSON(&alteration)
	err := alteration.Create(conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alteration)
}

func GetAlterations(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	flatId, _ := c.GetQuery("flat_id")

	altertions, err := model.GetAlterations(conn, flatId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, altertions)
}

func DeleteAlteration(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	alterationId, _ := c.GetQuery("alteration_id")
	err := model.DeleteAlteration(conn, alterationId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Alteration " + alterationId: "Deleted correctly"})
}

func UpdateAlteration(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)

	alteration := model.Alteration{}
	c.ShouldBindJSON(&alteration)
	err := alteration.UpdateAlteration(conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alteration)
}
