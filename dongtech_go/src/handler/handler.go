package handler

import (
	"dongtech_go/config"
	"dongtech_go/database"
	"dongtech_go/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

//http://localhost:9090/getUser/1
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, "参数错误")
		return
	}
	db := database.GetDB()
	var User struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	err = db.Table("users").Where("id = ?", id).Scan(&User).Error
	if err != nil {
		logrus.WithError(err).WithField("userId", id).Println("get err")
		util.Catch(err)
	}
	c.JSON(http.StatusOK, User)
}

//http://localhost:9090/version
func Version(c *gin.Context) {
	c.String(http.StatusOK, "V0.01")
}

//http://localhost:9090/index
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": fmt.Sprintf("%v,%v", "CreateByDongTech", time.Now().Local()),
	})
}

//http://localhost:9090/config
func PrintConfig(c *gin.Context) {
	config, err := config.GetConfig()
	if err != nil {
		logrus.WithError(err).Println("get config err")
		util.Catch(err)
	}
	c.String(http.StatusOK, fmt.Sprintf("%s,%d", config.Base.Author, config.Base.Age))
}

func UUID(c *gin.Context) {
	uuid := util.CreateUUID()
	c.String(http.StatusOK, uuid)
}
