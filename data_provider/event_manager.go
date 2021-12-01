package data_provider

import "sync"

type EventManager struct {
	mx sync.Mutex
	subsMatches map[int][] chan SerieData
	matches *MatchesStorage
}

func NewEventManager(matches *MatchesStorage) *EventManager {
	e := &EventManager{
		matches: matches,
	}
	e.subsMatches = make(map[int][]chan SerieData)

	return e
}

func (e *EventManager) SubscribeMatch(idMatch int, ch chan SerieData) {
	e.mx.Lock()
	defer e.mx.Unlock()
	e.subsMatches[idMatch] = append(e.subsMatches[idMatch], ch)
}

func (e *EventManager) PublishMatch(idMatch int, match SerieData) {
	e.mx.Lock()
	defer e.mx.Unlock()
	e.matches.Store(idMatch, match)

	for _, ch := range e.subsMatches[idMatch] {
		ch <- match
	}
}