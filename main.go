package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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

	r.Run(":9000")
}
