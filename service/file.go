package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type FileService struct {
}

func (h FileService) upload(c *gin.Context) {
	forms, err := c.MultipartForm()
	if err != nil {
		fmt.Println("error", err)
	}
	files := forms.File["fileName"]
	for _, v := range files {
		if err := c.SaveUploadedFile(v, fmt.Sprintf("%s%s", "./file/", v.Filename)); err != nil {
			fmt.Println("保存文件失败")
		}
	}
	return
}
