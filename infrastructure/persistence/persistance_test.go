package persistence_test

import (
	"database/sql"
	"os"
	"testing"
	"time"

	migrate "github.com/rubenv/sql-migrate"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
	"github.com/yasszu/go-firebase-auth-server/util/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestMain(m *testing.M) {
	txdb.Register("txdb", "postgres", conf.Postgres.TestDSN())

	txDB, err := sql.Open("txdb", conf.Postgres.TestDSN())
	if err != nil {
		panic(err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "../../migrations",
	}

	_, err = migrate.Exec(txDB, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}

func openTestDB() *gorm.DB {
	txDB, err := sql.Open("txdb", conf.Postgres.TestDSN())
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: txDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return gormDB
}

func now() time.Time {
	return time.Date(2021, 7, 5, 12, 13, 24, 0, time.UTC)
}
