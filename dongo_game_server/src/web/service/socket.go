package service

import (
	"dongo_game_server/src/database"
	"dongo_game_server/src/global_const"
	"dongo_game_server/src/model"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/zld126126/dongo_utils"
)

type SocketService struct {
	DB *database.DB
}

func (p *SocketService) NewSocketConfig(port int64) *model.SocketConfig {
	t := dongo_utils.Tick64()
	return &model.SocketConfig{
		Port:      port,
		ProjectId: 0,
		Status:    model.SocketStatusStop,
		Ct:        t,
		Mt:        t,
	}
}

func (p *SocketService) GetByPort(port int64) (*model.SocketConfig, error) {
	var c model.SocketConfig
	err := p.DB.Gorm.Table("socket_configs sc").Where(`sc.port = ?`, port).Scan(&c).Error
	if err != nil {
		if p.DB.IsGormNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (p *SocketService) Add(c *model.SocketConfig) error {
	err := p.DB.Gorm.Create(&c).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *SocketService) Save(c *model.SocketConfig) error {
	err := p.DB.Gorm.Save(&c).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *SocketService) GetInUsePort(projectId int64) (int, error) {
	type Result struct {
		Port int
	}
	var res Result
	err := p.DB.Gorm.Table("socket_configs sc").
		Where(`sc.project_id = ?`, projectId).
		Where(`sc.status = ?`, model.SocketStatusRun).
		Select("sc.port").
		Scan(&res).
		Error
	if err != nil {
		if p.DB.IsGormNotFound(err) {
			return 0, nil
		}
		return 0, err
	}
	return res.Port, nil
}

func (p *SocketService) GetAvailablePort(projectId int64) (int64, error) {
	type Result struct {
		Port int64
	}
	var res Result
	err := p.DB.Gorm.Table("socket_configs sc").
		Where(`sc.project_id = 0`).
		Where(`sc.status = ?`, model.SocketStatusStop).
		Order("sc.port asc").
		Select("sc.port").
		Limit(1).
		Scan(&res).
		Error
	if err != nil {
		return 0, err
	}
	return res.Port, nil
}

func (p *SocketService) UsePort(port int64, projectId int64) error {
	config, err := p.GetByPort(port)
	if err != nil {
		return err
	}

	config.ProjectId = projectId
	err = p.Save(config)
	if err != nil {
		return err
	}

	return nil
}

func (p *SocketService) InitPort() {
	for i := global_const.SocketPortMin; i < global_const.SocketPortMax; i++ {
		c, err := p.GetByPort(i)
		if err != nil {
			logrus.WithField("err", fmt.Sprintf("%+v", err)).WithField("port", i).Errorln("get socket config error")
			continue
		}

		if c != nil {
			c.Mt = dongo_utils.Tick64()
			c.Status = model.SocketStatusStop
			c.ProjectId = 0
			err = p.Save(c)
			if err != nil {
				logrus.WithField("err", fmt.Sprintf("%+v", err)).WithField("port", i).Errorln("create socket config error")
				continue
			}
			continue
		}

		c = p.NewSocketConfig(i)
		err = p.Add(c)
		if err != nil {
			logrus.WithField("err", fmt.Sprintf("%+v", err)).WithField("port", i).Errorln("create socket config error")
			continue
		}
	}
	logrus.Println("socket init success")
}

type ClientMessage struct {
	Token   string         `json:"token"`
	Project *model.Project `json:"project"`
	*model.Track
}

func (p *SocketService) DealMessage(s string) {

}
