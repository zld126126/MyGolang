package service

import (
	"dongo_game_server/src/database"
	"dongo_game_server/src/model"
	"errors"
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/zld126126/dongo_utils"
)

type ClientService struct {
	DB *database.DB
}

func (p *ClientService) CollectBySocket() {

}

func (p *ClientService) CollectByHttp(openId string, tp model.SourceType, token string, messages []string) {
	project, err := p.GetProject(token)
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`GetProject err`)
		return
	}

	c, err := p.GetConsumer(openId, tp)
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`GetConsumer err`)
		return
	}

	if c.Id == 0 {
		err := p.SaveConsumer(c)
		if err != nil {
			logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`SaveConsumer err`)
			return
		}
	}

	track, err := p.SaveTrack(project, c, tp, messages)
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`SaveTrack err`)
		return
	}

	logrus.WithField("track", track).Println("CollectByHttp success")
}

func (p *ClientService) GetConsumer(openId string, tp model.SourceType) (*model.Consumer, error) {
	var c model.Consumer
	err := p.DB.Gorm.Table(`consumers c`).
		Select(`c.*`).
		Joins(`left join consumer_items ci on ci.consumer_id = c.id`).
		Where(`ci.open_id = ?`, openId).Where(`ci.tp = ?`, tp).Scan(&c).Error
	if err != nil {
		if p.DB.IsGormNotFound(err) {
			return p.NewConsumer(openId, tp), nil
		}
		return nil, err
	}

	err = p.PrepareConsumerItem(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (p *ClientService) PrepareConsumerItem(c *model.Consumer) error {
	items := []*model.ConsumerItem{}
	err := p.DB.Gorm.Table("consumer_items ci").
		Where(`ci.consumer_id = ?`, c.Id).
		Order(`ci.id asc`).
		Scan(&items).Error
	if err != nil {
		return err
	}
	c.Items = items
	return nil
}

func (p *ClientService) NewConsumer(openId string, tp model.SourceType) *model.Consumer {
	dongo_utils.GenerateCode()
	t := dongo_utils.Tick64()
	items := []*model.ConsumerItem{}
	items = append(items, &model.ConsumerItem{
		Id:         0,
		ConsumerId: 0,
		Tp:         tp,
		OpenId:     openId,
		Ct:         t,
	})
	m := &model.Consumer{
		Id:    0,
		Ct:    t,
		Items: items,
	}
	return m
}

func (p *ClientService) SaveConsumer(c *model.Consumer) error {
	err := p.DB.Gorm.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&c).Error
		if err != nil {
			return err
		}

		for _, i := range c.Items {
			i.ConsumerId = c.Id
			err := tx.Create(&i).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *ClientService) GetProject(token string) (*model.Project, error) {
	var a model.Project
	err := p.DB.Gorm.Table(`projects p`).
		Where(`p.token = ?`, token).
		Scan(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (p *ClientService) NewTrack(projectId int64, consumerId int64, consumerItemId int64, tp model.SourceType, messages []string) *model.Track {
	tick := dongo_utils.Tick64()

	items := []*model.TrackItem{}
	for index, m := range messages {
		i := &model.TrackItem{
			Id:             0,
			TrackId:        0,
			ProjectId:      projectId,
			ConsumerId:     consumerId,
			ConsumerItemId: consumerItemId,
			SourceTp:       tp,
			Ct:             tick,
			MessageIndex:   int64(index),
			Message:        m,
		}
		items = append(items, i)
	}

	t := &model.Track{
		Id:             0,
		ProjectId:      projectId,
		ConsumerId:     consumerId,
		ConsumerItemId: consumerItemId,
		SourceTp:       tp,
		Ct:             tick,
		Items:          items,
		Messages:       messages,
	}
	return t
}

func (p *ClientService) SaveTrack(project *model.Project, consumer *model.Consumer, tp model.SourceType, messages []string) (*model.Track, error) {
	getCurrentConsumerItem := func() *model.ConsumerItem {
		item := linq.From(consumer.Items).WhereT(func(i *model.ConsumerItem) bool {
			return i.Tp == tp
		}).First().(*model.ConsumerItem)
		return item
	}

	item := getCurrentConsumerItem()
	if item == nil {
		return nil, errors.New("SaveTrack error")
	}

	track := p.NewTrack(project.Id, consumer.Id, item.Id, tp, messages)

	// 开启事务
	err := p.DB.Gorm.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&track).Error
		if err != nil {
			return err
		}

		for _, i := range track.Items {
			i.TrackId = track.Id
			err := tx.Create(&i).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return track, nil
}

// TODO 未来实现
func (p *ClientService) CollectByRpc() {

}
