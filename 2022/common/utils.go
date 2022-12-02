package common

import (
	"log"
	"os"
)

const PATH string = "../../input/2022/"

func LoadInput(env string, day string) string {
	data, err := os.ReadFile(PATH + env + "/" + day + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
