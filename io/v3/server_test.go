package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

type FileSystemStore struct {
	database io.ReadSeeker
}

func (f *FileSystemStore) GetLeague() []Player{
	f.database.Seek(0,0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemStore) GetPlayerScore(name string) int{
	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name{
			wins = player.Wins
			break
		}
	}
	return wins
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
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		store := FileSystemStore{database}

		got := store.GetPlayerScore("Chris")

		want := 33

		if got != want{
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func assertLeague(t *testing.T, got, want []Player){
	t.Helper()
	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v, want %v", got, want)
	}
}