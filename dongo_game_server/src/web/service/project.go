package service

import (
	"dongo_game_server/src/database"
	"dongo_game_server/src/global_const"
	"dongo_game_server/src/model"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/zld126126/dongo_utils/dongo_utils"
)

type ProjectService struct {
	DB *database.DB
}

func (p *ProjectService) Get(id int64) (*model.Project, error) {
	var a model.Project
	err := p.DB.Gorm.Table(`projects p`).
		Where(`p.dt = 0`).
		Where(`p.id = ?`, id).
		Scan(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func NewProject(name string, resourcePath string, restApi string, port int64) *model.Project {
	m := &model.Project{
		Name:         name,
		ResourcePath: resourcePath,
		RestApi:      restApi,
		Port:         port,
	}
	t := dongo_utils.Tick64()
	m.Ct = t
	m.Mt = t
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

func (p *ProjectService) GenerateToken(id int64, name string) string {
	return dongo_utils.GetMd5Str(global_const.ProjectTokenSalt)(fmt.Sprint(id) + name + fmt.Sprint(dongo_utils.Tick64()))
}

func (p *ProjectService) Add(name string, resourcePath string, restApi string, port int64) error {
	exist, err := p.ChkExist(name)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("不可创建同名用户")
	}

	m := NewProject(name, resourcePath, restApi, port)
	err = p.DB.Gorm.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&m).Error
		if err != nil {
			return err
		}

		m.Token = p.GenerateToken(m.Id, m.Name)
		err = tx.Save(&m).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectService) Del(id int64) error {
	m, err := p.Get(id)
	if err != nil {
		return err
	}

	m.Dt = dongo_utils.Tick64()
	err = p.DB.Gorm.Save(&m).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProjectService) List(name string, page int, pageSize int) (int, []*model.Project, error) {
	projects := []*model.Project{}
	total := 0
	_db := p.DB.Gorm.Table(`projects p`).
		Where(`p.dt = 0`)
	if name != "" {
		_db = _db.Where(`p.name like ?`, name)
	}
	err := _db.Count(&total).Error
	if err != nil {
		return 0, projects, err
	}

	if page > 0 && pageSize > 0 {
		_db = _db.Limit(pageSize).Offset(page - 1)
	}
	err = _db.Order("p.id desc").Find(&projects).Error
	if err != nil {
		return 0, projects, err
	}
	return total, projects, nil
}

func (p *ProjectService) RefreshToken(id int64) error {
	project, err := p.Get(id)
	if err != nil {
		return err
	}

	project.Token = p.GenerateToken(project.Id, project.Name)
	err = p.DB.Gorm.Save(&project).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectService) ChkExistForUpdate(name string, id int64) (bool, error) {
	total := 0
	err := p.DB.Gorm.Table(`projects p`).
		Where(`p.name = ?`, name).
		Where(`p.id != ?`, id).
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

func (p *ProjectService) Update(id int64, name string, resourcePath string, restApi string, port int64) error {
	exist, err := p.ChkExistForUpdate(name, id)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("用户名不可用")
	}

	m, err := p.Get(id)
	if err != nil {
		return err
	}

	m.Name = name
	m.ResourcePath = resourcePath
	m.RestApi = restApi
	m.Port = port
	m.Mt = dongo_utils.Tick64()

	err = p.DB.Gorm.Save(&m).Error
	if err != nil {
		return err
	}
	return nil
}
