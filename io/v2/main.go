package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore{
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string){
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player{
	var league []Player
	for name, wins := range i.store{
		league = append(league, Player{name, wins})
	}
	return league
}

func main(){
	server := NewPlayerServer(NewInMemoryPlayerStore())
	if err := http.ListenAndServe(":5000", server); err != nil{
		log.Fatalf("could not listen on port 5000, %v", err)
	}
}