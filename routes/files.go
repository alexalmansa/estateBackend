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
	filepath2 "path/filepath"
	"strconv"
	"strings"
)

func Upload(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	//Get flat from id
	flatId, _ := c.GetQuery("flat_id")
	flatModel := model.Flat{}

	flat, _ := flatModel.GetFlat(conn, flatId)
	if flat.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	file, header, err := c.Request.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	buildingId := strconv.Itoa(flat.BuildingId)
	filename := header.Filename

	nameNoSpaces := strings.ReplaceAll(filename, " ", "_")

	path := filepath2.Join("files/flats/building"+buildingId, "/flat"+flatId)

	filepath := path + "/" + nameNoSpaces

	if _, err := os.Stat(filepath); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "A file with the name : " + filename + " already exists"})
		return
	}
	err2 := os.MkdirAll(path, os.ModePerm)

	out, err := os.Create(filepath)
	if err != nil && err2 != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	error := saveFile(c, filename)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
	} else {
		c.JSON(http.StatusOK, gin.H{"File uploaded": filename})
	}
}

func saveFile(c *gin.Context, filepath string) error {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	file := model.Files{}

	flatId, _ := c.GetQuery("flat_id")
	i, err := strconv.Atoi(flatId)
	if err != nil {
		// handle error
		return err
	} else {
		return file.FileCreate(conn, i, filepath)
	}
}

func FilesFromFlat(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	flatId, _ := c.GetQuery("flat_id")
	file := model.Files{}

	flatFiles, err := file.GetFilesFromFlat(conn, flatId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, flatFiles)
}
func DownloadFile(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*sql.DB)
	flatId, _ := c.GetQuery("flat_id")
	fileName, _ := c.GetQuery("file_name")
	//Get flat from id
	flatModel := model.Flat{}
	flat, _ := flatModel.GetFlat(conn, flatId)
	if flat.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	path := filepath2.Join("./files/flats/building"+strconv.Itoa(flat.BuildingId), "/flat"+flatId)

	c.FileAttachment(path+"/"+fileName, fileName)
}
