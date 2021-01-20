package routes

import (
	"database/sql"
	"estateBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	nameNoSpaces := strings.ReplaceAll(filename, " ", "_")

	out, err := os.Create("public/building" + buildingId + "_flat" + flatId + "_" + nameNoSpaces)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:3000/file/building" + buildingId + "_flat" + flatId + "_" + nameNoSpaces
	error := saveFile(c, filepath)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
	} else {
		c.JSON(http.StatusOK, gin.H{"filepath": filepath})
	}
}

func saveFile(c *gin.Context, filepath string) error {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)

	flatId, _ := c.GetQuery("flat_id")
	i, err := strconv.Atoi(flatId)
	if err != nil {
		// handle error
		return err
	} else {
		return model.FileCreate(conn, i, filepath)
	}
}

func FilesFromFlat(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	flatId, _ := c.GetQuery("flat_id")
	flats, err := model.GetFilesFromFlat(conn, flatId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"files": flats})
}
