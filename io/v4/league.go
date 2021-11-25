package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(rdr io.Reader) ([]Player, error){
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil{
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league,err
}

type League []Player

func (l League) Find(name string) *Player{
	for i, p := range l {
		if p.Name == name{
			return &l[i]
		}
	}
	return nil
}

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league League
}

func NewFileSystemPlayerStore(databse io.ReadWriteSeeker) *FileSystemPlayerStore{
	databse.Seek(0,0)
	league, _ := NewLeague(databse)
	return &FileSystemPlayerStore{
		database: databse,
		league: league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() League{
	//f.database.Seek(0,0)
	//league, _ := NewLeague(f.database)
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int{
	//var wins int

	//for _, player := range f.GetLeague() {
	//	if player.Name == name{
	//		wins = player.Wins
	//		break
	//	}
	//}

	// player := f.GetLeague().Find(name)
	player := f.league.Find(name)

	if player != nil{
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string){
	// league := f.GetLeague()
	// player := league.Find(name)
	player := f.league.Find(name)

	//for i, player := range league{
	//	if player.Name == name{
	//		league[i].Wins++
	//	}
	//}

	if player != nil{
		player.Wins++
	}else{
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Seek(0,0)
	json.NewEncoder(f.database).Encode(f.league)
}