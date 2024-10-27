package po

import (
	"gorm.io/gorm"

	"gcms/internal/domain"
)

type User struct {
	gorm.Model

	Username string `gorm:"not null;default:'';unique;size:255;common:用户名"`
	Password string `gorm:"not null;default:'';size:128;comment:用户密码"`
	NickName string `gorm:"not null;default:'';size:64;comment:昵称"`
	Age      int    `gorm:"not null;default:0;comment:年龄"`
}

func (u User) TableName() string {
	return "t_gcms_user"
}

func ConvertToPO(do *domain.User) *User {
	return &User{
		Username: do.UserName,
		Password: do.Password,
		Age:      do.Age,
	}
}

func ConvertToDO(po *User) *domain.User {
	return &domain.User{
		ID:       po.ID,
		UserName: po.Username,
		Password: po.Password,
		NickName: po.NickName,
		Age:      po.Age,
	}
}
