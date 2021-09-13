package controller

import (
	"dongo_game_server/src/database"
	"dongo_game_server/src/global_const"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type BaseHdl struct {
	DB *database.DB
}

// @Summary 获取版本
// @Tags 获取版本
// @Description 获取版本
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /base/version [get]
// curl -X GET "127.0.0.1:9090/base/version"
func (p *BaseHdl) GetVersion() gin.HandlerFunc {
	version := viper.GetViper().GetString(global_const.ConfigVersionKey)
	return func(c *gin.Context) {
		c.Status(http.StatusOK)

		templateText := fmt.Sprintf("%v : %v", "GameServer_"+version, time.Now().Local())
		tmpl, err := template.New("version").Parse(templateText)
		if err != nil {
			log.Fatalf("parsing: %s", err)
		}
		tmpl.Execute(c.Writer, nil)
	}
}
