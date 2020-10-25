package main

import (
	"github.com/AyokunlePaul/crud-pay-api/src/api/application"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	if viperReadError := viper.ReadInConfig(); viperReadError != nil {
		panic(viperReadError)
	}
	application.StartApplication()
}
