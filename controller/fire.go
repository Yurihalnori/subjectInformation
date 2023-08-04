package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FileController struct {
}

func (s FileController) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	if _, err := os.Stat("./file/upload"); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll("./file/upload", os.ModePerm)
		}
	}
	c.SaveUploadedFile(file, "./file/upload/"+file.Filename)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    file,
	})
}