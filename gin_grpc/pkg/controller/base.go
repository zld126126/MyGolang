package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"gin_grpc/pkg/model"
	"gin_grpc/util"
)

// http://localhost:9090/user/1
func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "参数错误")
			return
		}
		var user model.User
		err = db.Table("users").Where("id = ?", id).Scan(&user).Error
		if err != nil {
			logrus.WithError(err).WithField("userId", id).Println("get err")
			util.Catch(err)
		}
		c.JSON(http.StatusOK, user)
	}
}

// http://localhost:9090/index
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": fmt.Sprintf("%v,%v", "CreateByDongTech", time.Now().Local()),
	})
}

func UUID(c *gin.Context) {
	uuid := util.CreateUUID()
	c.String(http.StatusOK, uuid)
}
