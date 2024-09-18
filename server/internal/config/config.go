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
		LogLevel:          viper.GetString("LogConfig.Level"),
		LogFormat:         viper.GetString("LogConfig.Format"),
		LogPath:           viper.GetString("LogConfig.Path"),
		LogName:           viper.GetString("LogConfig.Name"),
		LogFileMaxSize:    viper.GetInt("LogConfig.FileMaxSize"),
		LogFileMaxBackups: viper.GetInt("LogConfig.FileMaxBackups"),
		LogFileMaxAge:     viper.GetInt("LogConfig.FileMaxAge"),
		LogCompress:       viper.GetBool("LogConfig.Compress"),
		LogStdout:         viper.GetBool("LogConfig.Stdout"),
	}
}

func GetServerProtocol() string {
	return viper.GetString("ServerConfig.Prot")
}

func GetServerAddress() string {
	return viper.GetString("ServerConfig.Addr")
}
