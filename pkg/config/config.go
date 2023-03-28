package config

type Config struct {
	DBName   string
	DBDriver string
  Address string
}

func NewConfig() Config {
	return Config{
		DBName:   "./test.db",
		DBDriver: "sqlite3",
    Address : ":8080",
	}
}
