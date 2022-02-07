package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var imgsArr []string

type Imgs struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getImgsList() {
	var FileInfo []os.FileInfo
	var err error
	relativePath := "./files/imgs"

	if FileInfo, err = ioutil.ReadDir(relativePath); err != nil {
		fmt.Println("读取文件夹失败")
		return
	}

	baseUrl := "http://localhost:9000/static/"

	for _, fileInfo := range FileInfo {
		imgsObj := Imgs{"", baseUrl + fileInfo.Name()}
		data, err := json.Marshal(&imgsObj)
		if err != nil {
			fmt.Println("构造体转化JSON失败")
		}
		imgsArr = append(imgsArr, string(data))
	}
}

func main() {
	// log.Fatal(http.ListenAndServe(":8089", http.FileServer(http.Dir("files/imgs"))))

	r := gin.Default()
	r.StaticFS("/static", http.Dir("./files/imgs"))
	r.GET("/", func(c *gin.Context) {
		getImgsList()
		c.JSON(200, gin.H{
			"message": "pong",
			"imgsArr": imgsArr,
		})
	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})

	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Hello")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("username:%s, password:%s, type:%s", username, password, types))
	})

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}

		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})

	r.MaxMultipartMemory = 8 << 20

	r.POST("/mulupload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}

		files := form.File["files"]
		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}

		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})

	r.Run(":9000")
}
