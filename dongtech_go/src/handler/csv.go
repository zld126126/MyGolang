package handler

import (
	"bufio"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

const (
	FileMaxBytes = 1024 * 1024 * 2
)

//本地上传到服务器 csv格式，其他类似 读取内容
func UploadFile(c *gin.Context) {
	rFile, err := c.FormFile("file")
	if err != nil {
		c.String(400, "文件格式错误")
		return
	}

	if rFile.Size > FileMaxBytes {
		c.String(400, "文件大小超过2M")
		return
	}

	file, err := rFile.Open()
	if err != nil {
		c.String(400, "文件格式错误")
		return
	}
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()
		if err != nil {
			c.String(400, err.Error())
			return
		}
		//line 就是每一行的内容
		logrus.Println(line)
		//line[0] 就是第几列
		logrus.Println(line[0])
	}

}

//下载文件 读取内容
func DownloadReadFile(c *gin.Context) {
	//http下载地址 csv
	csvFileUrl := c.PostForm("file_name")
	res, err := http.Get(csvFileUrl)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	defer res.Body.Close()
	//读取csv
	reader := csv.NewReader(bufio.NewReader(res.Body))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			c.String(400, err.Error())
			return
		}
		//line 就是每一行的内容
		logrus.Println(line)
		//line[0] 就是第几列
		logrus.Println(line[0])
	}
}

//下载文件 写内容
func DownloadWriteFile(c *gin.Context) {
	//写文件
	var filename = "./output1.csv"
	if !checkFileIsExist(filename) {
		file, err := os.Create(filename) //创建文件
		if err != nil {
			c.String(400, err.Error())
			return
		}
		buf := bufio.NewWriter(file) //创建新的 Writer 对象
		buf.WriteString("test")
		buf.Flush()
		defer file.Close()
	}
	//返回文件流
	c.File(filename)
}

//判断文件是否存在  存在返回 true 不存在返回false
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
