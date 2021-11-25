package main

import (
	"fmt"
	"net/http"
)

//func PlayerServer(w http.ResponseWriter, r *http.Request){
//	player := r.URL.Path[len("/players/"):]
//
//	fmt.Fprintf(w, GetPlayerScore(player))
//}

func (p PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	player := r.URL.Path[len("/players/"):]

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

// PlayerServer
// 为了让 PlayerServer 能够使用 PlayerStore
type PlayerServer struct{
	store PlayerStore
}

type StubPlayerStore struct{
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

type PlayerStore interface {
	// GetPlayerScore 将重构的函数改为接口
	GetPlayerScore(name string) int
}


// GetPlayerScore 获取 Player 的分数
// 将分数检索分离来简化 PlayerServer
func GetPlayerScore(name string) string {
	if name == "Pepper"{
		return "20"
	}

	if name == "Floyd"{
		return "10"
	}

	return ""
}