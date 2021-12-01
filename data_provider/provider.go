package data_provider

type Provider struct {
	subsMatches map[int][] chan Match
}

func NewProvider () *Provider {
	//init listen
	return &Provider{}
}

func (p *Provider) listenMatches() {
	//subscriber := redisClient.Subscribe(ctx, "send-user-data")
	//
	//user := User{}
	//
	//for {
	//	msg, err := subscriber.ReceiveMessage(ctx)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
	//		panic(err)
	//	}
	//
	//	// ...
	//}
}

