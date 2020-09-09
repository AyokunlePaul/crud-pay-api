package main

import (
	"github.com/AyokunlePaul/crud-pay-api/src/api/application"
	"github.com/joho/godotenv"
)

func main() {
	environmentVariablesError := godotenv.Load(".env")
	if environmentVariablesError != nil {
		panic(environmentVariablesError)
	}
	application.StartApplication()
}
