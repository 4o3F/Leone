package utils

import (
	"github.com/spf13/viper"
	"log"
)

var (
	Monitormode bool
	Attackmode  bool
	Target      string
	Data        []string
)

func LoadConfig() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("./")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	Monitormode = config.GetBool("monitormode")
	Attackmode = config.GetBool("attackmode")
	Target = config.GetString("target")

	loadData()
}

func loadData() {
	data := viper.New()
	data.SetConfigName("data")
	data.AddConfigPath("./")
	err := data.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	Data = data.GetStringSlice("players")
}

func SaveData()  {
	data := viper.New()
	data.SetConfigName("data")
	data.SetConfigType("yml")
	data.AddConfigPath("./")
	err := data.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	data.Set("players", Data)
	err = data.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}
}
