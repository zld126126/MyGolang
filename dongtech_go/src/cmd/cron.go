package cmd

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func simpleCron() {
	for {
		fmt.Println("------this is a simple cron job", time.Now())
		time.Sleep(time.Second * 30)
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
