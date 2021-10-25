package data_provider

import "sync"

type Matches struct {
	mx sync.RWMutex
	m map[int]Match
}

type Match struct {
	Name string
}

func NewMatches() (*Matches, error) {
	m := &Matches{}

	return m, nil
}

func (m *Matches) Load(idMatch int) (Match, bool) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	val, ok := m.m[idMatch]

	return val, ok
}

func (m *Matches) Store(idMatch int, match Match){
	m.mx.Lock()
	defer m.mx.Unlock()
	m.m[idMatch] = match
}

