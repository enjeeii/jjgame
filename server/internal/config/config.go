package config

import (
	"github.com/spf13/viper"
	"jjgame/internal/logger"
	"log"
)

func InitAppConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("jjgame.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Panic("Init App Config Error:", err)
	}
}

func GetLoggerConfig() *logger.LogConfig {
	return &logger.LogConfig{
		LogLevel:          viper.GetString("Logger.Level"),
		LogFormat:         viper.GetString("Logger.Format"),
		LogPath:           viper.GetString("Logger.Path"),
		LogName:           viper.GetString("Logger.Name"),
		LogFileMaxSize:    viper.GetInt("Logger.FileMaxSize"),
		LogFileMaxBackups: viper.GetInt("Logger.FileMaxBackups"),
		LogFileMaxAge:     viper.GetInt("Logger.FileMaxAge"),
		LogCompress:       viper.GetBool("Logger.Compress"),
		LogStdout:         viper.GetBool("Logger.Stdout"),
	}
}

func GetString(name string) string {
	return viper.GetString(name)
}
