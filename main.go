package main

import (
	"os"

	"github.com/hash167/log-service/log"
)

func main() {
	f, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	_, err := log.NewStore(f)
	if err != nil {
		panic(err)
	}

}
