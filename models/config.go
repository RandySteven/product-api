package models

// dbName := os.Getenv("DB_NAME")
// dbPort := os.Getenv("DB_PORT")
// dbPass := os.Getenv("DB_PASS")
// dbHost := os.Getenv("DB_HOST")
// dbUser := os.Getenv("DB_USER")

type Config struct {
	DbName string
	DbPort string
	DbPass string
	DbHost string
	DbUser string
}

func NewConfig(dbHost, dbPort, dbUser, dbPass, dbName string) *Config {
	return &Config{
		DbName: dbName,
		DbPort: dbPort,
		DbPass: dbPass,
		DbHost: dbHost,
		DbUser: dbUser,
	}
}
