package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	flatId, _ := c.GetQuery("flat_id")
	buildingId, _ := c.GetQuery("building_id")
	/*if _, err := os.Stat("public/" + buildingId); os.IsNotExist(err) {
		os.Mkdir("public/" + buildingId, 7777)
	}*/
	out, err := os.Create("public/building" + buildingId + "_flat" + flatId + "_" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:3000/file/building" + buildingId + "_flat" + flatId + "_" + filename
	saveFile(c, filepath)

	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

func saveFile(c *gin.Context, filepath string) {
	/*db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	flatId, _ := c.GetQuery("flat_id")
	files, err := model.FileCreate(&conn, flatId, filepath)*/
}
