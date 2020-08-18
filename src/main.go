package main

import (
	"context"
	"estateBackend/model"
	"estateBackend/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
	"strings"
)

func main() {

	conn, err := connectDB()
	if err != nil {

	}
	r := gin.Default()
	r.Use(dbMiddleware(*conn))
	//Login endpoints
	usersGroup := r.Group("users")
	{
		usersGroup.POST("register", routes.UsersRegister)
		usersGroup.POST("login", routes.UsersLogin)
	}

	//Flats endpoints
	flatsGroup := r.Group("flats", authMiddleWare())
	{
		flatsGroup.GET("frombuilding", routes.FlatFromBuilding)
		flatsGroup.POST("create", routes.FlatCreate)
	}

	//Building endpoints
	buildingGroup := r.Group("buildings", authMiddleWare())
	{
		buildingGroup.POST("create", routes.BuildingCreate)
		buildingGroup.GET("getBuilding", routes.GetBuildings)

	}

	//Renters endpoints
	rentersGroup := r.Group("renters", authMiddleWare())
	{
		rentersGroup.POST("create", routes.RenterCreate)
	}

	//Lease endpoints
	leaseGroup := r.Group("leases", authMiddleWare())
	{
		leaseGroup.POST("create", routes.LeaseCreate)
		leaseGroup.GET("get", routes.GetLeases)

	}

	r.Run(":3000")
}

func connectDB() (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://alexalmansa:5554@localhost:5432/estate")
	if err != nil {
		fmt.Println("Error connecting to DB")
		fmt.Println(err.Error())
	}
	_ = conn.Ping(context.Background())
	return conn, err
}

func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := model.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
