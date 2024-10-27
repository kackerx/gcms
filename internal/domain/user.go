package domain

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"gcms/pkg/encrypt"
)

type User struct {
	ID       uint
	UserName string
	Password string
	NickName string
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
func (d *UserDomainService) GetUserByName(ctx context.Context, userName string) (u *User, exist bool, err error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetUserByName")
	defer func() {
		span.SetTag("req", userName)
		span.SetTag("resp", u)
		span.Finish()
	}()
	u, exist, err = d.repo.FindByName(userName)
	return u, exist, err
}
