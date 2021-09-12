package routes

import (
	"database/sql"
	model2 "estateBackend/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LeaseCreate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)

	lease := model2.Lease{}
	c.ShouldBindJSON(&lease)
	err := lease.Create(conn)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lease)
}

func GetLeases(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	flatId, _ := c.GetQuery("flat_id")
	renterid, _ := c.GetQuery("renter_id")
	pastLeases := "true"
	pastLeases, _ = c.GetQuery("past_leases")
	lease := model2.Lease{}

	leases, err := lease.GetAllLeases(conn, flatId, renterid, pastLeases)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, leases)
}

func DeleteLease(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	leaseId, _ := c.GetQuery("lease_id")
	lease := model2.Lease{}

	err := lease.DeleteLease(conn, leaseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Lease " + leaseId: "Deleted correctly"})
}

func UpdateLease(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)

	lease := model2.Lease{}
	c.ShouldBindJSON(&lease)
	err := lease.UpdateLease(conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lease)
}
