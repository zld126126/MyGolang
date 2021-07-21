package boot

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func simpleCron() {
	for {
		time.Sleep(time.Second * 5)
		fmt.Println("------this is a simple cron job", time.Now())
	}
}

func complexCron(c *cron.Cron) {
	cycle := `*/5 * * * * ?`
	cycleHourly := `@hourly`
	c.AddFunc(cycle, getTime)
	c.AddFunc(cycleHourly, getTime)
}

func getTime() {
	fmt.Println("======this is a complex cron job", time.Now())
}
