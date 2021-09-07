package service

import (
	"dongo_game_server/src/database"
	"dongo_game_server/src/global_const"
	"dongo_game_server/src/model"
	"dongo_game_server/src/util"

	"github.com/sirupsen/logrus"
)

type SocketService struct {
	DB *database.DB
}

func (p *SocketService) NewSocketConfig(port int) *model.SocketConfig {
	t := util.Tick64()
	return &model.SocketConfig{
		Port:      port,
		ProjectId: 0,
		Status:    model.SocketStatus_Stop,
		Ct:        t,
		Mt:        t,
	}
}

func (p *SocketService) GetByPort(port int) (*model.SocketConfig, error) {
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

func (p *SocketService) GetInUsePort(projectId int) (int, error) {
	type Result struct {
		Port int
	}
	var res Result
	err := p.DB.Gorm.Table("socket_configs sc").
		Where(`sc.project_id = ?`, projectId).
		Where(`sc.status = ?`, model.SocketStatus_Run).
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

func (p *SocketService) GetAvailablePort(projectId int) (int, error) {
	type Result struct {
		Port int
	}
	var res Result
	err := p.DB.Gorm.Table("socket_configs sc").
		Where(`sc.project_id = 0`).
		Where(`sc.status = ?`, model.SocketStatus_Stop).
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

func (p *SocketService) UsePort(port int, projectId int) error {
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
			logrus.WithField("port", i).WithError(err).Println("get socket config error")
			continue
		}

		if c != nil {
			c.Mt = util.Tick64()
			c.Status = model.SocketStatus_Stop
			c.ProjectId = 0
			err = p.Save(c)
			if err != nil {
				logrus.WithField("port", i).WithError(err).Println("create socket config error")
				continue
			}
			continue
		}

		c = p.NewSocketConfig(i)
		err = p.Add(c)
		if err != nil {
			logrus.WithField("port", i).WithError(err).Println("create socket config error")
			continue
		}
	}
	logrus.Println("socket init success")
}
