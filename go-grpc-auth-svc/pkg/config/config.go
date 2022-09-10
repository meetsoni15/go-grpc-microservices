package config

import (
	"log"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	DBUrl        string `mapstructure:"DB_URL"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	viperIns := viper.New()

	if env := os.Getenv("ENVRIONMENT"); env == "" {
		viperIns.AddConfigPath("./pkg/config/envs")
		viperIns.SetConfigName("dev")
		viperIns.SetConfigType("env")
		viperIns.AutomaticEnv()
		err = viperIns.ReadInConfig()
		if err != nil {
			return
		}
	} else {
		viperAutoMaticEnv(viperIns)
		viperIns.AutomaticEnv()
	}

	err = viperIns.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}

func viperAutoMaticEnv(viperIns *viper.Viper) {
	v := reflect.ValueOf(Config{})
	typeOfStruct := v.Type()

	for i := 0; i < v.NumField(); i++ {
		viperIns.BindEnv(typeOfStruct.Field(i).Tag.Get("mapstructure"))
		log.Println(typeOfStruct.Field(i).Tag.Get("mapstructure"))
	}

}
