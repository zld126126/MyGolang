package service

import (
	"dongo_game_server/src/database"
	"dongo_game_server/src/model"
)

type ProjectService struct {
	DB *database.DB
}

func (p *ProjectService) UseSocket(port int, projectId int) error {
	proj, err := p.GetByPort(port)
	if err != nil {
		return err
	}

	proj.Id = projectId
	err = p.DB.Gorm.Table(`projects p`).Save(&proj).Error
	if err != nil {
		return err
	}
	return nil
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

func (p *ProjectService) GetByPort(port int) (*model.Project, error) {
	var proj model.Project
	err := p.DB.Gorm.Table(`projects p`).
		Where(`p.dt = 0`).
		Where(`p.port = ?`, port).
		Scan(&proj).Error
	if err != nil {
		return nil, err
	}
	return &proj, nil
}

func (p *ProjectService) Create() error {

	return nil
}
