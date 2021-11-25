package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PlayerServer struct {
	store PlayerStore
	// 嵌入
	http.Handler
}
func NewPlayerServer(store PlayerStore) *PlayerServer{
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))

	p.Handler = router
	return p
}


type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	// GetPlayerScore 获取玩家得分
	GetPlayerScore(name string) int

	// RecordWin 记录玩家胜利次数
	RecordWin(name string)

	// GetLeague 获取联盟选手
	GetLeague() []Player
}

type StubPlayerStore struct {
	store map[string]int
	winCalls []string
}

func (s *StubPlayerStore) RecordWin(name string){
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetPlayerScore(name string) int{
	return s.store[name]
}



func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request){
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}
const jsonContentType = "application/json"

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	// JSON 编码
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string){
	score := p.store.GetPlayerScore(player)

	if score == 0{
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string){
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}