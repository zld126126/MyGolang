package cmd

import (
	"dongtech_go/config"
	"dongtech_go/proto"
	"dongtech_go/util"
	"fmt"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "DongTech",
	Short: "DongTech",
	Long:  `Example By DongBao`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logrus.Println("DongTech start")

		//启动 simple cron
		go simpleCron()

		//启动 cron
		c := cron.New()
		complexCron(c)
		c.Start()

		config, err := config.GetConfig()
		if err != nil {
			logrus.WithError(err).Println("get config err")
			util.Catch(err)
		}

		//启动 grpc
		go proto.CreateGrpcServe(fmt.Sprintf(":%s", config.Grpc.Addr))

		//启动web
		startWeb(config)

		logrus.Println("DongTech end")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
