package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Provider struct {
	subsMatches map[int][] chan Match
	EventManager *EventManager
}

var ctx = context.Background()
var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func NewProvider (eventManager *EventManager) *Provider {
	return &Provider{EventManager: eventManager}
}

func (p *Provider) listenMatches() {
	subscriber := redisClient.Subscribe(ctx, "live_score")
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		serieData := SerieData{}
		err = json.Unmarshal([]byte(msg.Payload), &serieData)
		fmt.Println(serieData)//////////
		if err != nil {
			panic(err)
		}
		p.EventManager.PublishSerie(1, serieData)
	}
}

