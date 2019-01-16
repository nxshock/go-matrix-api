package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	matrixapi "github.com/nxshock/go-matrix-api"
)

func init() {
	log.SetFlags(0)
}

func main() {
	client, err := matrixapi.NewClientWithPassword("http://127.0.0.1:8008", "userid", "password")
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Logout()

	client.InviteHandler = func(roomID, inviterID string) {
		err := client.JoinRoom(roomID)
		if err != nil {
			log.Println(err)
		}
	}

	client.MessageHandler = func(m matrixapi.IncomingMessage) {
		if m.SenderID != client.UserID {
			_, err := client.SendTextReply(m.RoomID, m.EventID, m.Body)
			if err != nil {
				log.Println(err)
			}
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-c:
			break
		case <-ticker.C:
			err = client.Sync()
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	_ = <-c
}
