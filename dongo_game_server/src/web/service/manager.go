package service

import (
	"dongo_game_server/src/database"
	"dongo_game_server/src/global_const"
	"dongo_game_server/src/model"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/zld126126/dongo_utils/dongo_utils"
)

type ManagerService struct {
	DB *database.DB
}

func NewManager(name string, passowrd string, tp model.ManagerType) *model.Manager {
	m := &model.Manager{
		Name:     name,
		Password: passowrd,
		Tp:       tp,
	}
	t := dongo_utils.Tick64()
	m.Ct = t
	m.Mt = t
	return m
}

func (p *ManagerService) ChkExist(name string) (bool, error) {
	total := 0
	err := p.DB.Gorm.Table(`managers m`).
		Where(`m.name = ?`, name).
		Where(`m.dt = 0`).
		Count(&total).Error
	if err != nil {
		if p.DB.IsGormNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return total > 0, nil
}

func (p *ManagerService) ChkExistForUpdate(name string, id int64) (bool, error) {
	total := 0
	err := p.DB.Gorm.Table(`managers m`).
		Where(`m.name = ?`, name).
		Where(`m.id != ?`, id).
		Where(`m.dt = 0`).
		Count(&total).Error
	if err != nil {
		if p.DB.IsGormNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return total > 0, nil
}

func (p *ManagerService) Login(name string, password string) (*model.Manager, error) {
	var user model.Manager
	err := p.DB.Gorm.Table(`managers m`).
		Where(`m.name = ?`, name).
		Where(`m.password = ?`, password).
		Where(`m.dt = 0`).
		Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *ManagerService) EncodeToken(m *model.Manager) (string, error) {
	if m == nil {
		return "", errors.New("错误的用户")
	}

	overTime := dongo_utils.Tick64(time.Now().AddDate(0, 0, 1))
	k := "%d" + global_const.ManagerLoginSplitKey + "%d"
	token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(k, m.Id, overTime)))
	return token, nil
}

func (p *ManagerService) DecodeToken(token string) (*model.Manager, error) {
	convert := func(s string) (int64, error) {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	}

	s, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, errors.New("错误的token")
	}

	arr := strings.Split(string(s), global_const.ManagerLoginSplitKey)
	if len(arr) != 2 {
		return nil, errors.New("错误的token")
	}

	managerId, err := convert(arr[0])
	if err != nil {
		return nil, errors.New("错误的用户信息")
	}

	m, err := p.Get(managerId)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	overTime, err := convert(arr[1])
	if err != nil {
		return nil, errors.New("用户信息已过期")
	}

	if overTime < dongo_utils.Tick64() {
		return nil, errors.New("用户信息已过期")
	}

	return m, nil
}

func (p *ManagerService) Get(id int64) (*model.Manager, error) {
	var user model.Manager
	err := p.DB.Gorm.Table(`managers m`).
		Where(`m.id = ?`, id).
		Where(`m.dt = 0`).
		Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *ManagerService) Add(name string, password string, tp model.ManagerType) error {
	exist, err := p.ChkExist(name)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("不可创建同名用户")
	}

	m := NewManager(name, password, tp)
	err = p.DB.Gorm.Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ManagerService) List(name string, page int, pageSize int) (int, []*model.Manager, error) {
	managers := []*model.Manager{}
	total := 0
	_db := p.DB.Gorm.Table(`managers m`).
		Where(`m.dt = 0`)
	if name != "" {
		_db = _db.Where(`m.name like ?`, name)
	}
	err := _db.Count(&total).Error
	if err != nil {
		return 0, managers, err
	}

	if page > 0 && pageSize > 0 {
		_db = _db.Limit(pageSize).Offset(page - 1)
	}
	err = _db.Order("m.id desc").Find(&managers).Error
	if err != nil {
		return 0, managers, err
	}
	return total, managers, nil
}

func (p *ManagerService) Update(id int64, name string, password string, tp model.ManagerType) error {
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
	m.Password = password
	m.Tp = tp
	m.Mt = dongo_utils.Tick64()

	err = p.DB.Gorm.Save(&m).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ManagerService) Del(id int64) error {
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
