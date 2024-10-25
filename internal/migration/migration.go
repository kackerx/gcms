package migration

import (
	"fmt"

	"gorm.io/gorm"

	"gcms/internal/data/po"
)

type Migration struct {
	db *gorm.DB
}

func NewMigration(db *gorm.DB) *Migration {
	return &Migration{db: db}
}

func (m *Migration) Run() {
	if err := m.db.AutoMigrate(&po.User{}); err != nil {
		panic(err)
	}

	fmt.Println("AutoMigrate end")
}
