package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		rw.WriteHeader(http.StatusOK)
	}))

	fasterServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fasterServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("Wanted '%s', but got '%s'", want, got)
	}

	slowServer.Close()
	fasterServer.Close()
}
