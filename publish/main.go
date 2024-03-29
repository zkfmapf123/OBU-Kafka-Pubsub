package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type OBUParams struct {
	OBUID int     `json:"obuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

var (
	sendInterval = 2
)

func main() {
	obuIds := genOUIDS(100)

	// kafka conn setting
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_HOST")})
	if err != nil {
		log.Fatalln(err)
	}

	defer p.Close()

	// check to producer events
	go func() {

		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed : %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered Success : %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "dkTopic"
	for i := 0; i < len(obuIds); i++ {
		lat, long := getLatLong()
		jsonr, err := json.Marshal(OBUParams{
			OBUID: obuIds[i],
			Lat:   lat,
			Long:  long,
		})

		if err != nil {
			log.Fatalln(err)
		}

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(jsonr),
		}, nil)

		time.Sleep(time.Second * time.Duration(sendInterval))
	}
	p.Flush(15 * 1000)
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
