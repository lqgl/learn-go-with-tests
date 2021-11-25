package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T){
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		// request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		// PlayerServer(response, request)
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		//if got != want{
		//	t.Errorf("got '%s', want '%s'",got , want)
		//}
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		// request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		// PlayerServer(response, request)
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		//if got != want{
		//	t.Errorf("got '%s', but want '%s'", got, want)
		//}
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		assertStatus(t, got, want)
	})
}

func newGetScoreRequest(name string) *http.Request{
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// 断言响应与预期结果是否一致
func assertResponseBody(t *testing.T, got, want string){
	t.Helper()
	if got != want{
		t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
	}
}

// 断言请求状态码
func assertStatus(t *testing.T, got, want int){
	t.Helper()
	if got != want{
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})
}