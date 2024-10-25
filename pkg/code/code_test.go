package code

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err := foo()

	var respErr *RespError
	if errors.As(err, &respErr) {
		fmt.Println(respErr)
	} else {
		fmt.Println("未知错误")
	}
}

func foo() error {
	return &RespError{ErrParameterInvalid, "参数错误"}
}
