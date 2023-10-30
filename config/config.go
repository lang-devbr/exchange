package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var (
	Port               int
	Env                string
	ReadTimeoutMillis  time.Duration
	WriteTimeoutMillis time.Duration
)

func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port = GetInt("HTTP_PORT")
	Env = GetString("ENV")
	ReadTimeoutMillis = GetDuration("HTTP_READ_TIMEOUT_MILLIS")
	WriteTimeoutMillis = GetDuration("HTTP_WRITE_TIMEOUT_MILLIS")
}

// GetString value of a given env var
func GetString(k string) string {
	return os.Getenv(k)
}

// GetInt value of a given env var
func GetInt(k string) int {
	v := GetString(k)
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return i
}

// GetDuration value of a given env var
func GetDuration(k string) time.Duration {
	return time.Duration(GetInt(k)) * time.Millisecond
}

// GetBool value of a given env var
func GetBool(k string) bool {
	v := GetString(k)
	return strings.ToLower(v) == "true"
}
