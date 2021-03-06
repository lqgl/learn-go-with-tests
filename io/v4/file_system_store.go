package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error){
	file.Seek(0,0)

	info, err := file.Stat()

	if err != nil{
		return nil, fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0{
		file.Write([]byte("[]"))
		file.Seek(0,0)
	}

	league, err := NewLeague(file)

	if err != nil{
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league: league,
	}, nil
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

	// f.database.Seek(0,0)
	//json.NewEncoder(f.database).Encode(f.league)
	f.database.Encode(f.league)
}
