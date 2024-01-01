package configs

var cfg *config

type config struct{
	API APIConfig
	DB DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Password string
	Database string
}

func InitConfig() {
	cfg = &config{
		API: APIConfig{
			Port: "8081",
		},
		DB: DBConfig{
			Host: "localhost",
			Port: "5432",
			User: "postgres",
			Password: "mrc03840",
			Database: "leadme",
		},
	}
}

func GetDB() DBConfig{
	return cfg.DB
}

func GetServerPort() string{
	return cfg.API.Port
}