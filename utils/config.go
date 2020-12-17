package utils

import (
	"github.com/spf13/viper"
	"log"
)

var (
	Monitormode bool
	Attackmode  bool
	Target      string
	Data        map[string]string
	//Player      []struct {
	//	Name string
	//	ID   uuid.UUID
	//}
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
	Data = data.GetStringMapString("players")
	//for i := range Data {
	//	name, id := CutData(Data[i])
	//	Player[i] = struct {
	//		Name string
	//		ID   uuid.UUID
	//	}{Name: name, ID: id}
	//}
}

func SaveData() {
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
