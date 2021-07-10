package db

import (
	"go-firebase-auth-server/util/conf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConn() (*gorm.DB, error) {
	return openConn()
}

func openConn() (*gorm.DB, error) {
	dialector := postgres.Open(conf.Postgres.DSN())
	return gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
