package env

type Variables struct {
	PostgresHost     Env
	PostgresPort     Env
	PostgresUser     Env
	PostgresPassword Env
	PostgresDB       Env
	PostgresTestDB   Env
	ServerHost       Env
	ServerPort       Env
}

func NewVariables() Variables {
	v := Variables{
		PostgresHost: Env{
			Key:   "POSTGRES_HOST",
			Value: "localhost",
		},
		PostgresPort: Env{
			Key:   "POSTGRES_PORT",
			Value: "5432",
		},
		PostgresUser: Env{
			Key:   "POSTGRES_USER",
			Value: "postgres",
		},
		PostgresPassword: Env{
			Key:   "POSTGRES_PASSWORD",
			Value: "password",
		},
		PostgresDB: Env{
			Key:   "POSTGRES_DB",
			Value: "postgres",
		},
		PostgresTestDB: Env{
			Key:   "POSTGRES_TEST_DB",
			Value: "postgres",
		},
		ServerHost: Env{
			Key:   "SERVER_HOST",
			Value: "localhost",
		},
		ServerPort: Env{
			Key:   "SERVER_PORT",
			Value: "8888",
		},
	}
	v.load()
	return v
}

func (v *Variables) load() {
	v.PostgresHost.Load()
	v.PostgresPort.Load()
	v.PostgresUser.Load()
	v.PostgresPassword.Load()
	v.PostgresDB.Load()
	v.PostgresTestDB.Load()
	v.ServerHost.Load()
	v.ServerPort.Load()
}
