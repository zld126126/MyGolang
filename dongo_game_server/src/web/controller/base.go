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