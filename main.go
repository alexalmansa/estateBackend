package main

import (
	"database/sql"
	"estateBackend/model"
	"estateBackend/routes"
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
		usersGroup.POST("register", routes.UsersRegister)
		usersGroup.POST("changePassword", routes.UsersChangePassword)
		usersGroup.POST("login", routes.UsersLogin)
		usersGroup.GET("me", routes.GetMe)
	}

	//Flats endpoints
	flatsGroup := Router.Group("flats", authMiddleWare())
	{
		flatsGroup.GET("frombuilding", routes.FlatFromBuilding)
		flatsGroup.POST("create", routes.FlatCreate)
		flatsGroup.DELETE("delete", routes.DeleteFlat)
		flatsGroup.PUT("edit", routes.UpdateFlat)

	}

	//Building endpoints
	buildingGroup := Router.Group("buildings", authMiddleWare())
	{
		buildingGroup.POST("create", routes.BuildingCreate)
		buildingGroup.GET("getBuilding", routes.GetBuildings)
		buildingGroup.DELETE("delete", routes.DeleteBuilding)
		buildingGroup.PUT("edit", routes.UpdateBuilding)
	}

	//Renters endpoints
	rentersGroup := Router.Group("renters", authMiddleWare())
	{
		rentersGroup.POST("create", routes.RenterCreate)
		rentersGroup.GET("getRenter", routes.GetRenter)
		rentersGroup.DELETE("delete", routes.DeleteRenter)
		rentersGroup.PUT("edit", routes.UpdateRenter)

	}

	//Alterations endpoints
	alterationsGroup := Router.Group("alterations", authMiddleWare())
	{
		alterationsGroup.POST("create", routes.AlterationCreate)
		alterationsGroup.GET("getAlterations", routes.GetAlterations)
		alterationsGroup.DELETE("delete", routes.DeleteAlteration)
		alterationsGroup.PUT("edit", routes.UpdateAlteration)
	}

	//Lease endpoints
	leaseGroup := Router.Group("leases", authMiddleWare())
	{
		leaseGroup.POST("create", routes.LeaseCreate)
		leaseGroup.GET("getLease", routes.GetLeases)
		leaseGroup.DELETE("delete", routes.DeleteLease)
		leaseGroup.PUT("edit", routes.UpdateLease)

	}

	//files endpoint
	filesGroup := Router.Group("files")
	{
		filesGroup.POST("/upload", routes.Upload)
		filesGroup.GET("getFiles", routes.FilesFromFlat)
		filesGroup.GET("downloadFile", routes.DownloadFile)
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
		user := model.User{}
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
