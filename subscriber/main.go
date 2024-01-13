package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type DataReceiver struct {
	msgCh chan interface{}
	conn  *websocket.Conn
}

func main() {
	recv := NewDataReceiver()
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgCh: make(chan interface{}, 128),
	}
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1028, 1028)
	if err != nil {
		log.Fatalln(err)
	}

	dr.conn = conn
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("Client Connect!!")

	for {
		var data interface{}

		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println(err)
			continue
		}

		dr.msgCh <- data
		go dr.wsReceiveLoop()
	}
}
