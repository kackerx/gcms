package handler

import (
	"github.com/gin-gonic/gin"

	"gcms/internal/service"
	"gcms/pkg/code"
	"gcms/pkg/help/resp"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	resp.ErrWithCode(c, code.ErrUnknownCode, "110")
}
func (h *UserHandler) Register(c *gin.Context) {
	var req *service.RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		resp.ErrWithCode(c, code.ErrParameterInvalid, err.Error())
		return
	}

	res, err := h.userService.Register(c, req)
	if err != nil {
		resp.Err(c, err)
		return
	}

	resp.Success(c, res)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req service.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		resp.ErrWithCode(c, code.ErrParameterInvalid, err.Error())
		return
	}

	login, err := h.userService.Login(c, &req)
	if err != nil {
		resp.Err(c, err)
		return
	}

	resp.Success(c, login)
}
