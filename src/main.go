package main

import (
	"context"
	"estateBackend/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)


func main() {

	conn, err := connectDB()
	if err != nil{

	}
	r := gin.Default()
	r.Use(dbMiddleware(*conn))
	usersGroup := r.Group("users")
	{
		usersGroup.POST("register", routes.UsersRegister)
	}
	r.Run(":3000")
}

func connectDB() (c *pgx.Conn, err error){
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