package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

type FileSystemStore struct {
	database io.Reader
}

func (f *FileSystemStore) GetLeague() []Player{
	//var league []Player
	//// JSON 解码
	//json.NewDecoder(f.database).Decode(&league)
	league, _ := NewLeague(f.database)
	return league
}

func TestFileSystemStore(t *testing.T){

	t.Run("/league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		store := FileSystemStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		// read again
		//got = store.GetLeague()
		//assertLeague(t, got, want)
	})
}

func assertLeague(t *testing.T, got, want []Player){
	t.Helper()
	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v, want %v", got, want)
	}
}