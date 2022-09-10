package config

import (
	"os"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	DBUrl         string `mapstructure:"DB_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
}

func LoadConfig() (config Config, err error) {
	viperIns := viper.New()

	if env := os.Getenv("ENVRIONMENT"); env == "" {
		viperIns.AddConfigPath("./pkg/config/envs")
		viperIns.SetConfigName("dev")
		viperIns.SetConfigType("env")
		viperIns.AutomaticEnv()
		if err = viperIns.ReadInConfig(); err != nil {
			return
		}
	} else {
		viperAutoMaticEnv(viperIns)
	}

	if err = viperIns.Unmarshal(&config); err != nil {
		return
	}

	return
}

func viperAutoMaticEnv(viperIns *viper.Viper) {
	v := reflect.ValueOf(Config{})
	typeOfStruct := v.Type()
	for i := 0; i < v.NumField(); i++ {
		viperIns.BindEnv(typeOfStruct.Field(i).Tag.Get("mapstructure"))
	}
}
