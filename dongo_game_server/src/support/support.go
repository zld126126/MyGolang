package support

import (
	"dongo_game_server/service/inf"
	config "dongo_game_server/src/config"
	"dongo_game_server/src/database"
	"fmt"
	"time"

	"github.com/robfig/cron"
)

type SupportApp struct {
	Config      *config.Config
	UserService inf.UserServiceClient
	DB          *database.DB
}

func (p *SupportApp) simpleCron() {
	for {
		fmt.Println("------this is a simple cron job", time.Now())
		time.Sleep(time.Second * 30)
	}
}

func (p *SupportApp) complexCron(c *cron.Cron) {
	cycle := `*/30 * * * * *`
	cycleHourly := `@hourly`
	c.AddFunc(cycle, p.getTime)
	c.AddFunc(cycleHourly, p.getTime)
}

func (p *SupportApp) getTime() {
	fmt.Println("======this is a complex cron job", time.Now())
}

func (p *SupportApp) Start() {
	go p.simpleCron()

	c := cron.New()
	go p.complexCron(c)
	c.Start()
}
