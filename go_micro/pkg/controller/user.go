package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"

	"go_micro/pkg/util"
	proto "go_micro/proto"
)

// http://localhost:9090/grpc/user/2
func GetGrpcUser(userService proto.QueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.String(http.StatusBadRequest, "参数错误")
			c.Abort()
			return
		}
		userResp, err := userService.GetUser(context.Background(), &proto.UserId{Id: int32(userId)})
		if err != nil {
			log.Fatalln(err)
			util.Catch(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"user_name": userResp.Name,
		})
	}
}
