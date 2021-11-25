package main

import (
	"fmt"
	"net/http"
)

type PlayerServer struct {
	store PlayerStore
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWins(name string)
}

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int{
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWins(name string){
	s.winCalls = append(s.winCalls, name)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := p.store.GetPlayerScore(player)

	if score == 0{
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string){
	w.WriteHeader(http.StatusAccepted)
	p.store.RecordWins(player)
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// ServeMux 多路请求复用器
	router := http.NewServeMux()
	// 将 http.Handler 附加到特定的请求路径上
	router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	}))

	router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		player := r.URL.Path[len("/players/"):]
		switch r.Method {
		case http.MethodGet:
			p.showScore(w, player)
		case http.MethodPost:
			p.processWin(w, player)
		}
	}))
}