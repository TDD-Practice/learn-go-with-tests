package racing

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowHttpServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(200 * time.Millisecond)
			w.WriteHeader(http.StatusOK)

		}))
	defer slowHttpServer.Close()

	fastHttpServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))
	defer fastHttpServer.Close()

	slowURL := slowHttpServer.URL
	fastURL := fastHttpServer.URL

	want := fastURL
	got, _ := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q but expected %q", got, want)
	}

	t.Run("Timeout at 10s",
		func(t *testing.T) {
			timeoutHttpServer := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					time.Sleep(11 * time.Second)
					w.WriteHeader(http.StatusOK)
				}))
			defer timeoutHttpServer.Close()

			_, err := Racer(slowURL, fastURL)

			if err == nil {
				t.Errorf("Expected to timeout after 10 sec")
			}
		})
}
