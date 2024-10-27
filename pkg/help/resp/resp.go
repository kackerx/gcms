package resp

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gcms/pkg/code"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	if data == nil {
		data = make(map[string]any)
	}
	c.JSON(http.StatusOK, response{0, "success", data})
}

func respError(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, response{code, msg, map[string]any{}})
}

func UnknownError(c *gin.Context, err error) {
	respError(c, code.ErrUnknownCode, err.Error())
}

func Err(c *gin.Context, err error) {
	var respErr *code.RespError
	if errors.As(err, &respErr) {
		respError(c, respErr.Code, respErr.Message)
		return
	}
	UnknownError(c, err)
}

func ErrWithCode(c *gin.Context, errCode int, args ...any) {
	e := code.ErrUnknown
	if v, ok := code.ErrMap[errCode]; ok {
		e = v
	}

	respError(c, e.Code, fmt.Sprintf(e.Message, args...))
}
