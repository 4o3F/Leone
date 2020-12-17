package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetAddr(address string) (string, int) {
	addr := strings.Split(address, ":")
	var port int
	switch len(addr) {
	case 1:
		port = 25565
	case 2:
		var err error
		port, err = strconv.Atoi(addr[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return addr[0], port
}
