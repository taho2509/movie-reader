package interfaces

import (
	"encoding/json"
	"fmt"
	"time"

	stan "github.com/nats-io/stan.go"
	uuid "github.com/satori/go.uuid"
)

var sc stan.Conn

func init() {
	var err error
	sc, err = stan.Connect("test-cluster", "csv-movie-reader", stan.NatsURL("nats://127.0.0.1:4222"))

	if err != nil {
		fmt.Println(err)

	}
}

type natsNotifier struct {
}

type Movie struct {
	Title string `json:"title"`
}

type Event struct {
	ID              string `json:"id"`
	Source          string `json:"source"`
	Specversion     string `json:"specversion"`
	Type            string `json:"type"`
	Datacontenttype string `json:"datacontenttype"`
	Subject         string `json:"subject"`
	Time            string `json:"time"`
	Data            Movie  `json:"data"`
}

func (*natsNotifier) Notify(title string) {
	u1 := uuid.NewV4()
	event := &Event{
		ID:              u1.String(),
		Source:          "/movie-reader-csv",
		Specversion:     "1.0",
		Type:            "com.movie-reader.movie.viewed",
		Datacontenttype: "application/json",
		Subject:         "movie_viewed",
		Time:            time.Now().Format(time.RFC3339),
		Data: Movie{
			Title: title,
		},
	}

	encodedEvent, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Simple Synchronous Publisher
	sc.Publish("movie_viewed", []byte(encodedEvent)) // does not return until an ack has been received from NATS Streaming
	fmt.Println(title)
}

// NewNatsNotifier return a concrete implementation of a DataNotifier
func NewNatsNotifier() *natsNotifier {
	return &natsNotifier{}
}
