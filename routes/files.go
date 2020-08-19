package routes

import (
	"estateBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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
	error := saveFile(c, filepath)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
	} else {
		c.JSON(http.StatusOK, gin.H{"filepath": filepath})
	}
}

func saveFile(c *gin.Context, filepath string) error {
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	flatId, _ := c.GetQuery("flat_id")
	i, err := strconv.Atoi(flatId)
	if err != nil {
		// handle error
		return err
	} else {
		return model.FileCreate(&conn, i, filepath)
	}
}
