package boot

import (
	"fmt"

	"github.com/robfig/cron"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "web demo",
	Long:  `web demo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("web cmd start")
		handle, cleanUp, err := InitHandle()
		if err != nil {
			panic(err)
		}
		defer cleanUp()
		handle.Run()
	},
}

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "cron demo",
	Long:  `cron demo`,
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			fmt.Println("cron cmd start")
			go simpleCron()
			c := cron.New()
			complexCron(c)
			c.Start()
		}()
	},
}

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "job demo",
	Long:  `job demo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("job cmd start")
	},
}

func init() {
	// aCmd.AddCommand(jobCmd)
}

func Execute() {
	cronCmd.Execute()
	webCmd.Execute()
	// 执行job
	// jobCmd.Execute()
}
