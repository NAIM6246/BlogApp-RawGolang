package configs

import "sync"

type DBConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

func NewDBConfig() *DBConfig {
	var loadDBOnce sync.Once
	loadDBOnce.Do(mapDBConfig)
	return dbConfig
}

var dbConfig *DBConfig

func mapDBConfig() {
	dbConfig = &DBConfig{
		Host:     "localhost",
		Port:     5432,
		DBName:   "go_assignment",
		User:     "naim",
		Password: "naim6246:)",
	}
}
