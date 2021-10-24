package data_provider

import "sync"

type Matches struct {
	mx sync.Mutex
	subs map[int][] chan Match
	m map[int]Match
}

type Match struct {
	Name string
}

func NewMatches() (*Matches, error) {
	m := &Matches{}
	m.subs = make(map[int][]chan Match)

	return m, nil
}

func (m *Matches) Load(idMatch int) (Match, bool) {
	m.mx.Lock()
	defer m.mx.Unlock()
	val, ok := m.m[idMatch]

	return val, ok
}

func (m *Matches) Store(idMatch int, match Match){
	m.mx.Lock()
	defer m.mx.Unlock()
	m.m[idMatch] = match
}

func (m *Matches) Subscribe(idMatch int, ch chan Match) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.subs[idMatch] = append(m.subs[idMatch], ch)
}

func (m *Matches) Publish(idMatch int, match Match) {
	m.Store(idMatch, match)
	m.mx.Lock()
	defer m.mx.Unlock()

	for _, ch := range m.subs[idMatch] {
		ch <- match
	}
}
