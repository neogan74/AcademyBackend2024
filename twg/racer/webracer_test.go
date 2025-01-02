package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one ", func(t *testing.T) {
		slowServer := makeDeleayedServer(20 * time.Millisecond)
		fastServer := makeDeleayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		if err != nil {
			t.Fatal("An error happens")
		}
	})

	t.Run("returns an error if a server doesnt respond withing 10s", func(t *testing.T) {
		serverA := makeDeleayedServer(11 * time.Second)
		serverB := makeDeleayedServer(21 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error after 10 sec but didn't get one")
		}
	})

}

func makeDeleayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
