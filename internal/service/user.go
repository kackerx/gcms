package service

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"gcms/internal/domain"
	"gcms/pkg/code"
)

type UserService struct {
	*Service
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

type UserDetailReq struct {
}

type UserDetailResp struct {
	NickName string
	Age      int
}

func NewUserService(svc *Service, ds *domain.UserDomainService) *UserService {
	return &UserService{Service: svc, ds: ds}
}

func (s *UserService) GetUser(ctx context.Context) (*UserDetailResp, error) {
	token, err := s.cache.Get(ctx, fmt.Sprintf(sessionKey, 1))
	if err != nil {
		return nil, err
	}

	parseToken, err := s.jwt.ParseToken(token)
	if err != nil {
		return nil, err
	}

	s.logger.Info("token", zap.String("uid", parseToken.UserId))
	return &UserDetailResp{}, nil
}

func (s *UserService) Login(ctx context.Context, req *LoginReq) (resp *LoginResp, err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("req", req)
	user, exist, err := s.ds.GetUserByName(ctx, req.Username)
	if err != nil {
		return
	}

	if !exist {
		return resp, code.NewErrWithCode(code.ErrUserNotExist, req.Username)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, code.NewErrWithCode(code.ErrUserPasswordInvalid)
	}

	token, err := s.jwt.GenToken(user.UserName, time.Now().Add(time.Hour))
	if err != nil {
		return
	}

	if err = s.cache.Set(ctx, fmt.Sprintf(sessionKey, user.ID), token, time.Hour); err != nil {
		return
	}

	return &LoginResp{
		SessionID: token,
		NickName:  user.NickName,
	}, nil
}

func (s *UserService) Register(ctx context.Context, req *RegisterReq) (resp *RegisterResp, err error) {
	_, exist, err := s.ds.GetUserByName(ctx, req.Username)
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
