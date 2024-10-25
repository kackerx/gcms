package data

import (
	"errors"
	"fmt"
	"log/slog"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"

	"gcms/internal/conf"
	"gcms/internal/data/po"
	"gcms/pkg/code"
	"gcms/vars"
)

type Data struct {
	db *gorm.DB
}

func NewDb(c *conf.Data) (*gorm.DB, func(), error) {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return db, func() {
		fmt.Println("db close")
	}, nil
}

func NewData(c *conf.Data) (data *Data, cleanup func(), err error) {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		return
	}

	return &Data{db: db}, func() {
		fmt.Println("close db")
	}, nil
}

func (d *Data) Create(userPo *po.User) (int, error) {
	if err := d.db.Create(userPo).Error; err != nil {
		// todo 记录数据库错误的日志, 不暴露给前端
		slog.Error(vars.ErrDBDesc, err.Error())
		return 0, code.ErrDB
	}
	return int(userPo.ID), nil
}

func (d *Data) SelectByCond(cond map[string]any) (*po.User, bool, error) {
	var user po.User
	err := d.db.Where(cond).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		slog.Error(vars.ErrDBDesc, err.Error())
		return nil, false, code.ErrDB
	}

	return &user, true, nil
}
