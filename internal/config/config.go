package config

import (
	"os"
)

// Config содержит параметры приложения
type Config struct {
    ARIURL      string
    ARIUsername string
    ARIPassword string
    ARIAppName  string
    DBPath      string
    HTTPPort    string
}

// LoadFromEnv загружает конфигурацию из переменных окружения
func LoadFromEnv() Config {
    cfg := Config{
        ARIURL:      getenv("ARI_URL", "http://192.168.88.205:8088/asterisk/ari"),
        ARIUsername: getenv("ARI_USERNAME", "asterisk_dev"),
        ARIPassword: getenv("ARI_PASSWORD", "32mS8SOEW0Cymrkr3SWaUUK4"),
        ARIAppName:  getenv("ARI_APP_NAME", "callcenter-app"),
        DBPath:      getenv("DB_PATH", "./callcenter.db"),
        HTTPPort:    getenv("HTTP_PORT", "3000"),
    }
    return cfg
}

func getenv(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}
