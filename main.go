package main

import (
	"fmt"
	"github.com/AkameMoe/Leone/component"
	"github.com/AkameMoe/Leone/utils"
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/google/uuid"
	"log"
)



func main()  {
	//component.LoadProxy()
	//log.Fatal(bot.OfflineUUID("CompoundKhan885"))
	go component.GetAddress()
	utils.LoadConfig()
	if utils.Monitormode {
		fmt.Println("Running in MonitorMode")
		component.StartMonitor()
	} else if utils.Attackmode {
		go component.StartMonitor()
		component.StartAttack()
	}



	//client := newJoin()
	//err := client.JoinServer("game.spawnmc.net",21117)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Login Success")
	//client.Events.ChatMsg = onChatMsg
	//client.Events.Disconnect = onDisconnect
	//err = client.HandleGame()
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func newJoin() *bot.Client {
	client := bot.NewClient()
	client.Name = "marsrockey"
	return client
}

func onChatMsg(chat chat.Message, pos byte, uuid uuid.UUID) error {
	log.Println("Chat:", chat)
	newJoin()
	return nil
}

func onDisconnect(reason chat.Message) error {
	log.Println(reason)
	return nil
}