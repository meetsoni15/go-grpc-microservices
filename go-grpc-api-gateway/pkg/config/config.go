package config

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSvcUrl   string `mapstructure:"ORDER_SVC_URL"`
}

func LoadConfig() (c Config, err error) {
	viperIns := viper.New()
	if env := os.Getenv("ENVRIONMENT"); env == "" {
		viperIns.AddConfigPath("./pkg/config/envs")
		viperIns.SetConfigName("dev")
		viperIns.SetConfigType("env")

		viperIns.AutomaticEnv()
		err = viperIns.ReadInConfig()
		if err != nil {
			fmt.Printf(`Config file not found because "%s"`, err)
			return
		}

	} else {
		viperAutoMaticEnv(viperIns)
	}

	err = viperIns.Unmarshal(&c)

	return
}

func viperAutoMaticEnv(viperIns *viper.Viper) {
	log.Println("Here")
	v := reflect.ValueOf(Config{})
	typeOfStruct := v.Type()

	for i := 0; i < v.NumField(); i++ {
		viperIns.BindEnv(typeOfStruct.Field(i).Tag.Get("mapstructure"))
	}

}
