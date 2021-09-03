package controller

import (
	"dongo_game_server/src/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ManagerHdl struct {
	DB *database.DB
}

// 登陆
func (p *ManagerHdl) Login(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	claims := &JWTClaims{
		UserID:      1,
		Username:    username,
		Password:    password,
		FullName:    username,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signedToken, err := getToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)
}

// 登出
func (p *ManagerHdl) LoginOut(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 验证
func (p *ManagerHdl) Verify(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 刷新令牌
func (p *ManagerHdl) Refresh(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
