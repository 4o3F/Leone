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
	"github.com/yezihack/colorlog"
)

var (
	proxyList []string
	proxyupdatedone bool
)

func LoadProxy() bot.Dialer {
	rand.Seed(time.Now().Unix())
	address := proxyList[rand.Intn(len(proxyList))]
	dia,err := proxy.SOCKS5("tcp", address, nil, proxy.Direct)
	if err != nil {
		colorlog.Warn("can't connect to the proxy:", err)
	}
	return dia
}

func GetAddress() {
	getAddress()
	proxylistcron := cron.New()
	proxylistcron.AddFunc("@every 10m", getAddress)
	proxylistcron.Start()

	select {}
}

func getAddress() {
	proxyupdatedone = false
	response, err := http.Get("https://www.proxy-list.download/api/v1/get?type=socks5")
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(body)))
	proxyList = proxyList[:0]
	for scanner.Scan() {
		colorlog.Info("Checking proxy " + scanner.Text())
		if checkProxy(scanner.Text()) {
			proxyList = append(proxyList, scanner.Text())
		}
	}
	colorlog.Info("Get Proxy List Success")
}

func checkProxy(address string) bool {
	dialer, err := proxy.SOCKS5("tcp", address, nil, proxy.Direct)
	if err != nil {
		return false
	}
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport, Timeout: 5*time.Second}
	httpTransport.Dial = dialer.Dial
	resp, err := httpClient.Get("https://www.baidu.com")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	colorlog.Warn("Proxy " + address + " avaliable")
	proxyupdatedone = true
	return true
}