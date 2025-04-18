package goatcounter

import (
	"log/slog"
	"os"
	"testing"
)

const ()

func testClient() *Client {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	instance := "goatcounter.com"

	code := os.Getenv("CODE")
	token := os.Getenv("TOKEN")

	return NewClient(code, instance, token, logger)
}

// TestCount tests GoatCounter /count endpoint
func TestCount(t *testing.T) {
	client := testClient()

	t.Run("count", func(t *testing.T) {
		count, err := client.Count()
		if err != nil {
			msg := "expected success"
			if errResponseor, ok := err.(*ErrorResponse); ok {
				t.Fatalf("%s\n%+v", msg, errResponseor)
			}
			t.Fatal(msg, err)
		}

		t.Log(count)
	})
}

// TestPaths tests GoatCounter Paths endpoints
func TestPaths(t *testing.T) {
	client := testClient()

	t.Run("list", func(t *testing.T) {
		paths, err := client.Paths().List()
		if err != nil {
			msg := "expected success"
			if errResponseor, ok := err.(*ErrorResponse); ok {
				t.Fatalf("%s\n%+v", msg, errResponseor)
			}
			t.Fatal(msg, err)
		}

		t.Log(paths)
	})
}

// TestSites tests GoatCounter Sites endpoints
func TestSites(t *testing.T) {
	client := testClient()

	t.Run("get", func(t *testing.T) {
		ID := "11407"
		site, err := client.Sites().Get(ID)
		if err != nil {
			msg := "expected success"
			if errResponseor, ok := err.(*ErrorResponse); ok {
				t.Fatalf("%s\n%+v", msg, errResponseor)
			}
			t.Fatal(msg, err)
		}

		t.Log(site)
	})
	t.Run("list", func(t *testing.T) {
		sites, err := client.Sites().List()
		if err != nil {
			msg := "expected success"
			if errResponseor, ok := err.(*ErrorResponse); ok {
				t.Fatalf("%s\n%+v", msg, errResponseor)
			}
			t.Fatal(msg, err)
		}

		t.Log(sites)
	})
}

// TestStats tests GoatCounter Statistics endpoints
func TestStats(t *testing.T) {
	client := testClient()

	t.Run("hits", func(t *testing.T) {
		hits, err := client.Stats().Hits()
		if err != nil {
			msg := "expected success"
			if errResponseor, ok := err.(*ErrorResponse); ok {
				t.Fatalf("%s\n%+v", msg, errResponseor)
			}
			t.Fatal(msg, err)
		}

		t.Log(hits)
	})
	t.Run("total", func(t *testing.T) {
		total, err := client.Stats().Total()
		if err != nil {
			msg := "expected success"
			if errResponseor, ok := err.(*ErrorResponse); ok {
				t.Fatalf("%s\n%+v", msg, errResponseor)
			}
			t.Fatal(msg, err)
		}

		t.Log(total)
	})
}

// TestUsers tests GoatCounter Users endpoints
func TestUsers(t *testing.T) {
	client := testClient()

	t.Run("me", func(t *testing.T) {
		user, err := client.Users().Me()
		if err != nil {
			msg := "expected success"
			if errResponseor, ok := err.(*ErrorResponse); ok {
				t.Fatalf("%s\n%+v", msg, errResponseor)
			}
			t.Fatal(msg, err)
		}

		t.Log(user)
	})
}
