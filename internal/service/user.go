package service

import (
	"context"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"gcms/internal/domain"
	"gcms/pkg/code"
)

type UserService struct {
	ds *domain.UserDomainService
}

type RegisterReq struct {
	Username string `json:"username" bind:"required"`
	Password string `json:"password" bind:"required"`
	Age      int    `json:"age"`
}

type RegisterResp struct {
	id int `json:"id"`
}

type LoginReq struct {
	Username string `json:"username" bind:"required"`
	Password string `json:"password" bind:"required"`
}

type LoginResp struct {
	SessionID string `json:"session_id"`
	NickName  string `json:"nick_name"`
}

func NewUserService(ds *domain.UserDomainService) *UserService {
	return &UserService{ds: ds}
}

func (s *UserService) GetUser(c *gin.Context) {
	return
}

func (s *UserService) Login(ctx context.Context, req *LoginReq) (resp *LoginResp, err error) {
	user, exist, err := s.ds.GetUserByName(req.Username)
	if err != nil {
		return
	}

	if !exist {
		return resp, code.NewErrWithCode(code.ErrUserNotExist, req.Username)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, code.NewErrWithCode(code.ErrUserPasswordInvalid)
	}

	return &LoginResp{
		SessionID: "110",
		NickName:  "",
	}, nil
}

func (s *UserService) Register(ctx context.Context, req *RegisterReq) (resp *RegisterResp, err error) {
	_, exist, err := s.ds.GetUserByName(req.Username)
	if err != nil {
		return
	}

	if exist {
		return resp, code.NewErrWithCode(code.ErrUserExist, req.Username)
	}

	uid, err := s.ds.CreateUser(&domain.User{
		UserName: req.Username,
		Password: req.Password,
		Age:      req.Age,
	})
	if err != nil {
		return
	}

	return &RegisterResp{id: uid}, nil
}
