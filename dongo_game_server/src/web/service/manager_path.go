package service

import (
	"dongo_game_server/src/database"
	"dongo_game_server/src/global_const"
	"dongo_game_server/src/model"
	"errors"
)

type ManagerPathService struct {
	DB *database.DB
}

func (p *ManagerPathService) ChkExist(optionPath string) (bool, error) {
	if global_const.IsNormalPath(optionPath) {
		return true, nil
	}

	total := 0
	err := p.DB.Gorm.Table(`manager_paths m`).
		Where(`m.option_path = ?`, optionPath).
		Count(&total).Error
	if err != nil {
		if p.DB.IsGormNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return total > 0, nil
}

func (p *ManagerPathService) ChkExistForUpdate(id int64, optionPath string) (bool, error) {
	if global_const.IsNormalPath(optionPath) {
		return true, nil
	}

	total := 0
	err := p.DB.Gorm.Table(`manager_paths m`).
		Where(`m.option_path = ?`, optionPath).
		Where(`m.id != ?`, id).
		Count(&total).Error
	if err != nil {
		if p.DB.IsGormNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return total > 0, nil
}

func NewManagerPath(name string, optionPath string) *model.ManagerPath {
	m := &model.ManagerPath{
		Name:       name,
		OptionPath: optionPath,
	}
	return m
}

func (p *ManagerPathService) Add(name string, optionPath string) error {
	exist, err := p.ChkExist(optionPath)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("不可创建同路径地址")
	}

	m := NewManagerPath(name, optionPath)
	err = p.DB.Gorm.Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ManagerPathService) List(optionPath string, page int, pageSize int) (int, []*model.ManagerPath, error) {
	paths := []*model.ManagerPath{}
	total := 0
	_db := p.DB.Gorm.Table(`manager_paths m`)
	if optionPath != "" {
		_db = _db.Where(`m.option_path like ?`, optionPath)
	}
	err := _db.Count(&total).Error
	if err != nil {
		return 0, paths, err
	}

	if page > 0 && pageSize > 0 {
		_db = _db.Limit(pageSize).Offset(page - 1)
	}
	err = _db.Order("m.id desc").Find(&paths).Error
	if err != nil {
		return 0, paths, err
	}
	return total, paths, nil
}

func (p *ManagerPathService) Get(id int64) (*model.ManagerPath, error) {
	var m model.ManagerPath
	err := p.DB.Gorm.Table(`manager_paths m`).
		Where(`m.id = ?`, id).
		Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (p *ManagerPathService) Update(id int64, name string, optionPath string) error {
	exist, err := p.ChkExistForUpdate(id, optionPath)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("地址不可用")
	}

	m, err := p.Get(id)
	if err != nil {
		return err
	}

	m.Name = name
	m.OptionPath = optionPath

	err = p.DB.Gorm.Save(&m).Error
	if err != nil {
		return err
	}
	return nil
}
