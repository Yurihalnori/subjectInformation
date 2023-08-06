package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FileController struct {
}

func (s FileController) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	table := c.PostForm("table")
	field := c.PostForm("field")
	dst := "./file/" + table + "/" + field + "/"
	fmt.Println(dst)
	os.MkdirAll(dst, os.ModePerm)
	c.SaveUploadedFile(file, dst+file.Filename)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    file,
	})
}

func (s FileController) Download(c *gin.Context) {

}
