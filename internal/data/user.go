package data

import (
	"gcms/internal/data/po"
	"gcms/internal/domain"
)

type UserRepo struct {
	data *Data
}

func (r *UserRepo) FindByName(name string) (*domain.User, bool, error) {
	userPO, exist, err := r.data.SelectByCond(map[string]any{"username": name})
	if err != nil {
		return nil, false, err
	}

	if !exist {
		return nil, false, nil
	}

	return po.ConvertToDO(userPO), exist, nil
}

func (r *UserRepo) SaveUser(user *domain.User) (int, error) {
	return r.data.Create(po.ConvertToPO(user))
}

func (r *UserRepo) FindByID(id int64) (*domain.User, error) {
	// TODO implement me
	panic("implement me")
}

func NewUserRepo(data *Data) domain.UserRepo {
	return &UserRepo{data: data}
}

func (r *UserRepo) Hello() string {
	return "hello world"
}
