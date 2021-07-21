package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"

	"gin_grpc/service/inf"
	"gin_grpc/util"
)

// http://localhost:9090/grpc/user/2
func GetGrpcUser(userService inf.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.String(http.StatusBadRequest, "参数错误")
			c.Abort()
			return
		}
		userResp, err := userService.GetUser(context.Background(), &inf.UserReq{Id: int32(userId)})
		if err != nil {
			log.Fatalln(err)
			util.Catch(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"user_name": userResp.Name,
		})
	}
}
