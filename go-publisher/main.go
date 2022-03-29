package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/zeromq/goczmq"
)

type Notification struct {
	Message   string
	Timestamp int64
}

func main() {
	pubUrl := os.Getenv("PUB_URL")
	if pubUrl == "" {
		pubUrl = "tcp://*:5555"
	}

	pub, err := goczmq.NewPub(pubUrl)
	exitOnError(err)
	defer pub.Destroy()

	fmt.Printf("Accepting connections at %v\n", pubUrl)

	time.Sleep(1 * time.Second)

	count, err := strconv.Atoi(os.Getenv("PUBLISH_COUNT"))
	if count < 1 || err != nil {
		count = 42
	}

	for i := 0; i < count; i++ {
		text := fmt.Sprintf("from go-publisher: %v", i)

		message, err := json.Marshal(&Notification{
			Message:   text,
			Timestamp: time.Now().Unix(),
		})
		exitOnError(err)

		fmt.Printf("Publishing: %v\n", text)
		err = pub.SendFrame(message, 0)
		exitOnError(err)

		time.Sleep(1 * time.Second)
	}
}

func exitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
