package cmd

import (
	"cobra-viper/src/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "DongTech",
	Short: "DongTech",
	Long:  `Example Cobra By DongBao`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logrus.Println("DongTech cobra start")
		//simple cron
		go simpleCron()

		c := cron.New()
		complexCron(c)
		c.Start()

		router := gin.Default()
		router.GET("/version", version)
		router.GET("/config", printConfig)
		router.Run(":9090")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func simpleCron() {
	for {
		fmt.Println("------this is a simple cron job", time.Now())
		time.Sleep(time.Second * 30)
	}
}

func version(c *gin.Context) {
	c.String(http.StatusOK, "V0.1")
}

func printConfig(c *gin.Context) {
	config, err := config.GetConfig()
	if err != nil {
		logrus.WithError(err).Println("get config err")
	} else {
		c.String(http.StatusOK, fmt.Sprintf("%s,%d", config.Base.Author, config.Base.Age))
	}
}

func complexCron(c *cron.Cron) {
	cycle := `*/30 * * * * *`
	cycleHourly := `@hourly`
	c.AddFunc(cycle, getTime)
	c.AddFunc(cycleHourly, getTime)
}

func getTime() {
	fmt.Println("======this is a complex cron job", time.Now())
}
