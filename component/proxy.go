package component

import (
	"bufio"
	"github.com/Tnze/go-mc/bot"
	"github.com/robfig/cron"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var proxyList []string

func LoadProxy() bot.Dialer {
	rand.Seed(time.Now().Unix())
	address := proxyList[rand.Intn(len(proxyList))]
	dia,err := proxy.SOCKS5("tcp", address, nil, proxy.Direct)
	if err != nil {
		log.Fatalf("can't connect to the proxy:", err)
	}
	return dia
}

func GetAddress() {
	getAddress()
	proxylistcron := cron.New()
	proxylistcron.AddFunc("@every 20s", getAddress)
	proxylistcron.Start()

	select {}
}

func getAddress() {
	response, err := http.Get("https://www.proxy-list.download/api/v1/get?type=socks5&country=CN")
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(body)))
	for scanner.Scan() {
		proxyList = append(proxyList, scanner.Text())
	}
	log.Println("Get Proxy List Success")
}