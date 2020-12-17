package utils

import (
	"github.com/google/uuid"
	"log"
	"strings"
)

func CutData(data string) (string, uuid.UUID) {
	str := strings.Split(data, ";")

	name := str[0]
	uuid, err := uuid.FromBytes([]byte(str[1]))

	if err != nil {
		log.Fatal(err)
	}
	return name, uuid
}
