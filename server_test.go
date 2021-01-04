package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const jsonContentType = "application/json"

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	tests := map[string]struct {
		player string
		want   string
	}{
		"returns Pepper's score": {"Pepper", "20"},
		"returns Floyd's score":  {"Floyd", "10"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			request := newGetScoreRequest(test.player)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertStatus(t, response.Code, http.StatusOK)
			assertResponseBody(t, response.Body.String(), test.want)
		})

	}
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		AssertPlayerWin(t, &store, player)
	})
}

func TestLeague(t *testing.T) {

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertContentType(t, response, jsonContentType)
		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
	})
}
