package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := viper.New()

	config.AddConfigPath(path)
	config.SetConfigName("ip")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	for _, part := range config.GetStringSlice("ports") {
		scan(part)
	}
	fmt.Println("scan done...")
}
