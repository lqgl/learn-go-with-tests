package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestFileSystemStore(t *testing.T){

	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTemplate(t,`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		// store := FileSystemPlayerStore{database}
		store := NewFileSystemPlayerStore(database)


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
		database, cleanDatabase := createTemplate(t,`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		//store := FileSystemPlayerStore{database}
		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Chris")

		want := 33

		if got != want{
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTemplate(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		// store := FileSystemPlayerStore{database}
		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")

		want := 1
		assertScoreEquals(t, got, want)
	})
}

func assertLeague(t *testing.T, got, want []Player){
	t.Helper()
	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v, want %v", got, want)
	}
}

func createTemplate(t *testing.T, initialData string) (io.ReadWriteSeeker, func()){
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil{
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertScoreEquals(t *testing.T, got, want int){
	t.Helper()
	if got != want{
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestRecordingAndRetrievingThem(t *testing.T){
	database, cleanDatabase := createTemplate(t, "")
	defer cleanDatabase()

	// store := &FileSystemPlayerStore{database}
	store := NewFileSystemPlayerStore(database)
	
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("got league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
}

func NewLeagueRequest() *http.Request{
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newPostWinRequest(name string) *http.Request{
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s",name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request{
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatus(t *testing.T, got, want int){
	t.Helper()
	if got != want{
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string){
	t.Helper()
	if got != want{
		t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
	}
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil{
		t.Errorf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}
	return
}