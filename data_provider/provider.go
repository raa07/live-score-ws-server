package data_provider

type Provider struct {
	subsMatches map[int][] chan Match
}

func NewProvider () *Provider {
	//init listen
	return &Provider{}
}

func (p *Provider) listenMatches() {
	///event manager publish match
}

