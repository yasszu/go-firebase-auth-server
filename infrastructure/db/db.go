package db

import (
	"database/sql"
	"log"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/yasszu/go-firebase-auth-server/util/conf"
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

func Migration(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	log.Printf("sql-migrate: Applied %d migrations!\n", n)

	return nil
}
