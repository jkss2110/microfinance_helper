package models

import (
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/websocket"
)

type message struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func LoanDashboardSocket(ws *websocket.Conn) {
	done := make(chan struct{})
	fmt.Println("Connection established")
	go func(c *websocket.Conn) {
		for {
			var msg message
			if err := websocket.JSON.Receive(ws, &msg); err != nil {
				log.Println(err)
				break
			}
			fmt.Printf("received message %s\n", msg.Data)
		}
		close(done)
	}(ws)
loop:
	for {
		select {
		case <-done:
			fmt.Println("Socket connection closed")
			break loop
		default:
			filter := bson.D{}
			// Sorting on Loan Amount to show in the dashboard
			sorter := bson.D{{"loanAmt", 1}}
			// setting a limit to 5 and needs to change it too
			var limit int64 = 5
			results, err := GetLoanDetailTopDB(filter, sorter, limit)
			if err != nil {
				log.Println(err)
				break
			}
			if err = websocket.JSON.Send(ws, results); err != nil {
				log.Println(err)
				break
			}
			time.Sleep(10 * time.Second)
		}
	}
	defer ws.Close()
	fmt.Println("Connection closed")
}
