package domain

import "gcms/pkg/encrypt"

type User struct {
	UserName string
	Password string
	Age      int
}

type UserDomainService struct {
	repo UserRepo
}

func NewUserDomainService(repo UserRepo) *UserDomainService {
	return &UserDomainService{repo: repo}
}

func (d *UserDomainService) GetUserByID(uid int64) (*User, error) {
	return d.repo.FindByID(uid)
}

func (d *UserDomainService) CreateUser(user *User) (int, error) {
	password, err := encrypt.EncryptPassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = password
	return d.repo.SaveUser(user)
}
func (d *UserDomainService) GetUserByName(userName string) (*User, bool, error) {
	return d.repo.FindByName(userName)
}
