package data_provider

import "sync"

type EventManager struct {
	mx sync.Mutex
	subsMatches map[int][] chan Match
	matches *Matches
}

func NewEventManager(matches *Matches) *EventManager {
	e := &EventManager{
		matches: matches,
	}
	e.subsMatches = make(map[int][]chan Match)

	return e
}

func (e *EventManager) SubscribeMatch(idMatch int, ch chan Match) {
	e.mx.Lock()
	defer e.mx.Unlock()
	e.subsMatches[idMatch] = append(e.subsMatches[idMatch], ch)
}

func (e *EventManager) PublishMatch(idMatch int, match Match) {
	e.mx.Lock()
	defer e.mx.Unlock()
	e.matches.Store(idMatch, match)

	for _, ch := range e.subsMatches[idMatch] {
		ch <- match
	}
}