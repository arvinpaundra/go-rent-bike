package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort              string `mapstructure:"APP_PORT"`
	DBUsername           string `mapstructure:"DBUSERNAME"`
	DBPassword           string `mapstructure:"DBPASSWORD"`
	DBAddress            string `mapstructure:"DBADDRESS"`
	DBName               string `mapstructure:"DBNAME"`
	JWTSecret            string `mapstructure:"JWT_SECRET"`
	MidtransServerKeyDev string `mapstructure:"MIDTRANS_SERVER_KEY_DEV"`
	AuthString           string `mapstructure:"AUTH_STRING"`
}

var Cfg *Config

func InitConfig() {
	cfg := &Config{}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error read env: %v", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		fmt.Printf("error while unmarshal env: %v", err)
	}

	Cfg = cfg
}
