package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

type OBUParams struct {
	OBUID int     `json:"obuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

var (
	wsEndPoint   = ""
	sendInterval = 2
)

func main() {
	wsEndPoint = fmt.Sprintf("ws://%s:%s/ws", os.Getenv("SUB_HOST"), os.Getenv("PORT"))
	obuIds := genOUIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndPoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for i := 0; i < len(obuIds); i++ {
			lat, long := getLatLong()

			data := OBUParams{
				OBUID: obuIds[i],
				Lat:   lat,
				Long:  long,
			}

			fmt.Printf("%v\n", data)
			if err := conn.WriteJSON((data)); err != nil {
				log.Fatalln(err)
			}

			time.Sleep(time.Second * time.Duration(sendInterval))
		}
	}
}

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func getLatLong() (lat float64, long float64) {
	return genCoord(), genCoord()
}

func genOUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
