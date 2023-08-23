package goatcounter

import (
	"log"
	"os"
	"testing"
)

const (
	testCode   string   = "example"
	testMethod Endpoint = Me
)

func TestEndpoint(t *testing.T) {
	want := "https://example.goatcounter.com/api/v0/me"

	client := NewClient(testCode, "")
	got := client.Url(testMethod)

	if got != want {
		t.Errorf("got:%s\nwant: %s", got, want)
	}
}

func TestMe(t *testing.T) {
	code := os.Getenv("CODE")
	token := os.Getenv("TOKEN")
	client := NewClient(code, token)
	user, err := client.Me()
	if err != nil {
		t.Fatal("expected success")
	}

	log.Println(user)
}

func TestStats(t *testing.T) {
	code := os.Getenv("CODE")
	token := os.Getenv("TOKEN")
	client := NewClient(code, token)

	t.Run("hits", func(t *testing.T) {
		hits, err := client.Stats().Hits()
		if err != nil {
			t.Fatal("expected success", err)
		}

		log.Println(hits)
	})
	t.Run("total", func(t *testing.T) {
		total, err := client.Stats().Total()
		if err != nil {
			t.Fatal("expected success", err)
		}

		log.Println(total)
	})
}
