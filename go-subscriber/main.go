package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/zeromq/goczmq"
)

type Notification struct {
	Message   string
	Timestamp int64
}

func main() {
	receiveTimeout := 10_000
	connectUrl := os.Getenv("CONNECT_URL")
	if connectUrl == "" {
		connectUrl = "tcp://localhost:5555"
	}
	sub, err := goczmq.NewSub(connectUrl, "" /*all*/)
	exitOnError(err)
	sub.SetRcvtimeo(receiveTimeout)
	defer sub.Destroy()

	log.Printf("connecting to %v", connectUrl)

	pid := os.Getpid()

	for {
		msg, _, err := sub.RecvFrame()
		exitOnError(err)
		message := Notification{}
		exitOnError(json.Unmarshal(msg, &message))
		log.Printf("%v received: %v, sent at %v", pid, message.Message, message.Timestamp)
	}
}

func exitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
