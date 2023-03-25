package config

type Config struct {
	DBName   string
	DBDriver string
}

func NewConfig() Config {
	return Config{
		DBName:   "./test.db",
		DBDriver: "sqlite3",
	}
}
