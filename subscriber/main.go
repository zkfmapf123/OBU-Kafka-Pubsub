package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type DataReceiver struct {
	msgCh chan interface{}
	conn  *websocket.Conn
}

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/ws", recv.handleWS)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err)
	}
}

func NewDataReceiver() (*DataReceiver, error) {
	return &DataReceiver{
		msgCh: make(chan interface{}, 128),
	}, nil
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
	}
	dr.conn = conn
	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("Client Connect!!")

	for {
		var data map[string]interface{}
		reqId := rand.Intn(math.MaxInt)
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println(err)
			continue
		}

		data["RequestId"] = reqId
		fmt.Println(data)
	}
}
