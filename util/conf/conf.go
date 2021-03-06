package conf

import (
	"fmt"

	"github.com/yasszu/go-firebase-auth-server/util/env"
)

var (
	Server   *server
	Postgres *postgres
)

func init() {
	v := env.NewVariables()
	Server = &server{
		Port: v.ServerPort.Value,
		Host: v.ServerHost.Value,
	}
	Postgres = &postgres{
		Host:     v.PostgresHost.Value,
		Port:     v.PostgresPort.Int(),
		Username: v.PostgresUser.Value,
		Password: v.PostgresPassword.Value,
		DB:       v.PostgresDB.Value,
		TestDB:   v.PostgresTestDB.Value,
	}
}

type server struct {
	Port string
	Host string
}

func (s *server) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
	TestDB   string
}

func (p *postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		p.Host,
		p.Port,
		p.Username,
		p.DB,
		p.Password)
}

func (p *postgres) TestDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		p.Host,
		p.Port,
		p.Username,
		p.TestDB,
		p.Password)
}
