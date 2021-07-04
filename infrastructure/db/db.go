package db

import (
	"fmt"
	"time"

	"go-firebase-auth-server/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	retryTimes  = 20
	waitingTime = 3 * time.Second
)

func NewConn() (*gorm.DB, error) {
	db, err := openConn()
	for i := 0; i < retryTimes; i++ {
		if err == nil {
			break
		}
		fmt.Println("Waiting for getting the connection of Postgres...")
		time.Sleep(waitingTime)
		db, err = openConn()
	}

	return db, err
}

func openConn() (*gorm.DB, error) {
	dialector := postgres.Open(util.NewConf().DSN())
	return gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
