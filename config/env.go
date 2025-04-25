package config

type (
	PostgresConfig struct {
		Timeout  int
		DBName   string
		Username string
		Password string
		Host     string
		Port     string
	}
	SqlConfig struct {
		Timeout  int
		DBName   string
		Username string
		Password string
		Host     string
		Port     string
	}



	

	EnvConfig struct {
		Host     string
		Port     int
		Sql      SqlConfig
		Postgres PostgresConfig
	}
)
