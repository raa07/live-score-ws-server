package main

import (
	"fmt"
	"sync"
)

type EventManager struct {
	mx         sync.Mutex
	subsSeries map[int][] chan<- SerieData
	series     *SeriesStorage
}

func NewEventManager(matches *SeriesStorage) *EventManager {
	e := &EventManager{
		series: matches,
	}
	e.subsSeries = make(map[int][]chan<- SerieData)

	return e
}

func (e *EventManager) SubscribeSerie(serieId int, ch chan<- SerieData) {
	e.mx.Lock()
	defer e.mx.Unlock()
	fmt.Println(ch)
	e.subsSeries[serieId] = append(e.subsSeries[serieId], ch)
}

func (e *EventManager) PublishSerie(serieId int, serie SerieData) {
	e.mx.Lock()
	defer e.mx.Unlock()
	e.series.Store(serieId, serie)
	fmt.Println(e.subsSeries)////

	for _, ch := range e.subsSeries[serieId] {
		fmt.Println(ch)
		ch <- serie
	}
}