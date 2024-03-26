package config

import (
	"fmt"
	"os"
	"time"

	"github.com/BoburjonRaximov/premier_league/storage"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Limit   int
	Page    int
	Methods []string
	Objects []string
}

const (
	SuccessStatus = iota + 1
	CancelStatus
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"

	TimeExpiredAt = time.Hour * 720
)

func Load() *Config {
	return &Config{
		Limit:   10,
		Page:    1,
		Methods: []string{"create", "get", "update","getAll", "delete"},
		Objects: []string{"clubs", "goals", "match"},
	}
}

type ConfigPostgres struct {
	Environment string // debug, test, release

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	Port string

	PostgresMaxConnections int32
}

// Club implements storage.StorageI.
func (ConfigPostgres) Club() storage.ClubesI {
	panic("unimplemented")
}

// Goal implements storage.StorageI.
func (ConfigPostgres) Goal() storage.GoalsI {
	panic("unimplemented")
}

// Match implements storage.StorageI.
func (ConfigPostgres) Match() storage.MatchesI {
	panic("unimplemented")
}

// Results implements storage.StorageI.
func (ConfigPostgres) Result() storage.ResultsI {
	panic("unimplemented")
}

// Load ...
func LoadP() ConfigPostgres {
	if err := godotenv.Load("./.env"); err != nil {
		fmt.Println("No .env file found")
	}

	config := ConfigPostgres{}

	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Port = cast.ToString(getOrReturnDefaultValue("PORT", "8080"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1234"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "league_db"))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}
