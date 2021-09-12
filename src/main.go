package main

import (
	"database/sql"
	model2 "estateBackend/src/model"
	routes2 "estateBackend/src/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strings"
)

// CORS Middleware
func CORS(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

func main() {

	conn, err := connectDB()
	if err != nil {

	}
	Router := gin.Default()
	Router.Use(dbMiddleware(conn))
	Router.Use(CORS)
	//Login endpoints
	usersGroup := Router.Group("users")
	{
		usersGroup.POST("register", routes2.UsersRegister)
		usersGroup.POST("changePassword", routes2.UsersChangePassword)
		usersGroup.POST("login", routes2.UsersLogin)
		usersGroup.GET("me", routes2.GetMe)
	}

	//Flats endpoints
	flatsGroup := Router.Group("flats", authMiddleWare())
	{
		flatsGroup.GET("frombuilding", routes2.FlatFromBuilding)
		flatsGroup.POST("create", routes2.FlatCreate)
		flatsGroup.DELETE("delete", routes2.DeleteFlat)
		flatsGroup.PUT("edit", routes2.UpdateFlat)

	}

	//Building endpoints
	buildingGroup := Router.Group("buildings", authMiddleWare())
	{
		buildingGroup.POST("create", routes2.BuildingCreate)
		buildingGroup.GET("getBuilding", routes2.GetBuildings)
		buildingGroup.DELETE("delete", routes2.DeleteBuilding)
		buildingGroup.PUT("edit", routes2.UpdateBuilding)
	}

	//Renters endpoints
	rentersGroup := Router.Group("renters", authMiddleWare())
	{
		rentersGroup.POST("create", routes2.RenterCreate)
		rentersGroup.GET("getRenter", routes2.GetRenter)
		rentersGroup.DELETE("delete", routes2.DeleteRenter)
		rentersGroup.PUT("edit", routes2.UpdateRenter)

	}

	//Alterations endpoints
	alterationsGroup := Router.Group("alterations", authMiddleWare())
	{
		alterationsGroup.POST("create", routes2.AlterationCreate)
		alterationsGroup.GET("getAlterations", routes2.GetAlterations)
		alterationsGroup.DELETE("delete", routes2.DeleteAlteration)
		alterationsGroup.PUT("edit", routes2.UpdateAlteration)
	}

	//Lease endpoints
	leaseGroup := Router.Group("leases", authMiddleWare())
	{
		leaseGroup.POST("create", routes2.LeaseCreate)
		leaseGroup.GET("getLease", routes2.GetLeases)
		leaseGroup.DELETE("delete", routes2.DeleteLease)
		leaseGroup.PUT("edit", routes2.UpdateLease)

	}

	//files endpoint
	filesGroup := Router.Group("files")
	{
		filesGroup.POST("/upload", routes2.Upload)
		filesGroup.GET("getFiles", routes2.FilesFromFlat)
		filesGroup.GET("downloadFile", routes2.DownloadFile)
	}
	Router.Static("/file", "./public")

	Router.Run(":3000")
}

func connectDB() (c *sql.DB, err error) {
	conn, err := sql.Open("mysql", "alexalmansa:test@tcp(localhost:3306)/estate")
	if err != nil {
		fmt.Println("Error connecting to DB")
		fmt.Println(err.Error())
	}

	err = conn.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	return conn, err
}

func dbMiddleware(conn *sql.DB) gin.HandlerFunc {
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
		user := model2.User{}
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := user.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
