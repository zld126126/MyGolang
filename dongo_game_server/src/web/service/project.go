package service

import (
	"dongo_game_server/src/database"
	"dongo_game_server/src/model"
	"errors"

	"github.com/zld126126/dongo_utils/dongo_utils"
)

type ProjectService struct {
	DB *database.DB
}

func (p *ProjectService) Get(id int) (*model.Project, error) {
	var proj model.Project
	err := p.DB.Gorm.Table(`projects p`).
		Where(`p.dt = 0`).
		Where(`p.id = ?`, id).
		Scan(&proj).Error
	if err != nil {
		return nil, err
	}
	return &proj, nil
}

func NewProject(name string, resourcePath string, restApi string) *model.Project {
	m := &model.Project{
		Name:         name,
		ResourcePath: resourcePath,
		RestApi:      restApi,
	}
	t := dongo_utils.Tick64()
	m.Ct = t
	m.Mt = t

	m.Token = dongo_utils.GetMd5Str("simple")(name)
	return m
}

func (p *ProjectService) ChkExist(name string) (bool, error) {
	total := 0
	err := p.DB.Gorm.Table(`projects p`).
		Where(`p.name = ?`, name).
		Where(`p.dt = 0`).
		Count(&total).Error
	if err != nil {
		if p.DB.IsGormNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return total > 0, nil
}

func (p *ProjectService) Add(name string, resourcePath string, restApi string) error {
	exist, err := p.ChkExist(name)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("不可创建同名用户")
	}

	m := NewProject(name, resourcePath, restApi)
	err = p.DB.Gorm.Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}
