package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

var index = 13321

func Producer(url string) {
	w := &kafka.Writer{
		Addr:     kafka.TCP("65.1.9.139:9092"),
		Topic:    "prod.trell_crawler_profile",
		Balancer: &kafka.LeastBytes{},
	}
	//print(url)
	errWriter := w.WriteMessages(context.Background(),
		kafka.Message{
			//Key:   []byte("Key-"+string(index)),
			Value: []byte("" + url),
		},
	)
	index = index + 1
	if errWriter != nil {
		log.Fatal("failed to write message:", errWriter)
	}
	if errClose := w.Close(); errClose != nil {
		log.Fatal("failsed to close writer:", errClose)
	}
}
func StartKafka() {

	conf := kafka.ReaderConfig{
		Brokers: []string{"65.1.9.139:9092"},
		Topic:   "prod.trell_crawler_profile",
		GroupID: "gi",
	}

	reader := kafka.NewReader(conf)
	fmt.Println("kafka-started")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Some error occured", err)
			continue
		}
		fmt.Println("Message is:", string((m.Value)))
	}
}
