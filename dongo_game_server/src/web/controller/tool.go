package controller

import (
	"bufio"
	"dongo_game_server/src/config"
	"dongo_game_server/src/database"
	"dongo_game_server/src/global_const"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"github.com/zld126126/dongo_utils/dongo_utils"
)

type ToolHdl struct {
	DB    *database.DB
	Email *config.EmailConfig
}

//本地上传到服务器 csv格式，其他类似 读取内容
func (p *ToolHdl) UploadFile(c *gin.Context) {
	rFile, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "文件格式错误")
		return
	}

	if rFile.Size > global_const.FileMaxBytes {
		c.String(http.StatusBadRequest, "文件大小超过2M")
		return
	}

	file, err := rFile.Open()
	if err != nil {
		c.String(http.StatusBadRequest, "文件格式错误")
		return
	}
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		//line 就是每一行的内容
		logrus.Println(line)
		//line[0] 就是第几列
		logrus.Println(line[0])
	}

}

//下载文件 读取内容
func (p *ToolHdl) DownloadReadFile(c *gin.Context) {
	//http下载地址 csv
	csvFileUrl := c.PostForm("file_name")
	res, err := http.Get(csvFileUrl)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
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
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		//line 就是每一行的内容
		logrus.Println(line)
		//line[0] 就是第几列
		logrus.Println(line[0])
	}
}

//下载文件 写内容
func (p *ToolHdl) DownloadWriteFile(c *gin.Context) {
	//写文件
	var filename = "./output1.csv"
	if !checkFileIsExist(filename) {
		file, err := os.Create(filename) //创建文件
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
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

// 发送邮件
func (p *ToolHdl) SendEmail(c *gin.Context) {
	var email dongo_utils.EmailConfig
	err := copier.Copy(&email, p.Email)
	if err != nil {
		logrus.WithError(err).Println("send email err")
		c.String(http.StatusBadRequest, "send email err")
		return
	}

	var msg dongo_utils.EmailMsg
	msg.From = []string{"zld126126@sina.com"}
	// 10分钟邮箱
	msg.To = []string{"QWYHIN@10min.club"}
	msg.Subject = fmt.Sprintf("%s,%d", "test", dongo_utils.ParseSecondTimeToInt64())
	msg.Body = "this is a test letter from dongbao."
	msg.From = []string{email.Username}

	err = dongo_utils.SendEmail(&msg, &email)
	if err != nil {
		logrus.WithError(err).Println("send email err")
		c.String(http.StatusBadRequest, "send email err")
		return
	}

	c.String(http.StatusOK, "send email ok")
}
