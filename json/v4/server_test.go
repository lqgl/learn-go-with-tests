package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLeague(t *testing.T){
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)
	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player
		// JSON 解码
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil{
			t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func assertStatus(t *testing.T, got, want int){
	t.Helper()
	if got != want{
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}