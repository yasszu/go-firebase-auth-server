package persistence_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"go-firebase-auth-server/util"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestMain(m *testing.M) {
	txdb.Register("txdb", "postgres", util.NewConf().DSN())
	code := m.Run()
	os.Exit(code)
}

func openTestDB() *gorm.DB {
	txDB, err := sql.Open("txdb", util.NewConf().TestDSN())
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: txDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	return gormDB
}
