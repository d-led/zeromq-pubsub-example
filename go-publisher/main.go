package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

	log.Printf("Accepting connections at %v", pubUrl)

	const count = 42
	for i := 0; i < count; i++ {
		text := fmt.Sprintf("Message: %v", i)

		message, err := json.Marshal(&Notification{
			Message:   text,
			Timestamp: time.Now().Unix(),
		})
		exitOnError(err)

		log.Printf("Publishing: %v", text)
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
