//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"gcms/internal/conf"
	"gcms/internal/data"
	"gcms/internal/migration"
)

var migrationSet = wire.NewSet(
	migration.NewMigration,
)

var databaseSet = wire.NewSet(
	data.NewDb,
)

func wireApp(data *conf.Data) (*migration.Migration, func(), error) {
	panic(wire.Build(databaseSet, migrationSet))
}
