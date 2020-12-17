package component

import (
	"github.com/AkameMoe/Leone/utils"
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/google/uuid"
	"github.com/json-iterator/go"
	"github.com/robfig/cron"
	"log"
)

type status struct {
	Description chat.Message
	Players     struct {
		Max    int
		Online int
		Sample []struct {
			ID   uuid.UUID
			Name string
		}
	}
	Version struct {
		Name     string
		Protocol int
	}
	//favicon ignored
}

func StartMonitor()  {
	monitorcron := cron.New()
	monitorcron.AddFunc("@every 5s", doMonitor)
	monitorcron.Start()

	savecron := cron.New()
	savecron.AddFunc("@every 30s", doSave)
	savecron.Start()
	select {}
}

func doMonitor() {
	addr, port := utils.GetAddr(utils.Target)
	response , _,err := bot.PingAndList(addr, port)
	if err != nil {
		log.Println(err)
	}
	var status status
	err = jsoniter.Unmarshal(response, &status)
	if err != nil {
		log.Println(err)
	}
	for i := range status.Players.Sample {
		exist := utils.InStringSlice(utils.Data, status.Players.Sample[i].Name)
		if !exist {
			utils.Data = append(utils.Data, status.Players.Sample[i].Name)
			log.Println("New Player: " + status.Players.Sample[i].Name)
		}
	}
	//log.Println("Monitor " + strconv.Itoa(int(time.Now().UnixNano()/1e6)))
}

func doSave()  {
	utils.SaveData()
	//log.Println("Save " + strconv.Itoa(int(time.Now().UnixNano()/1e6)))
}