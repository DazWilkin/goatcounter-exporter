package goatcounter

import (
	"log"
	"os"
	"testing"
)

const (
	testCode string = "example"
)

// TestCount tests GoatCounter /count endpoint
func TestCount(t *testing.T) {
	code := os.Getenv("CODE")
	token := os.Getenv("TOKEN")
	client := NewClient(code, token)

	t.Run("count", func(t *testing.T) {
		count, err := client.Count()
		if err != nil {
			msg := "expected success"
			if errResponseor, ok := err.(*ErrorResponse); ok {
				t.Fatalf("%s\n%+v", msg, errResponseor)
			}
			t.Fatal(msg, err)
		}

		log.Println(count)
	})
}

// TestPaths tests GoatCounter Paths endpoints
func TestPaths(t *testing.T) {
	code := os.Getenv("CODE")
	token := os.Getenv("TOKEN")
	client := NewClient(code, token)

	t.Run("list", func(t *testing.T) {
		paths, err := client.Paths().List()
		if err != nil {
			msg := "expected success"
			if errResponseor, ok := err.(*ErrorResponse); ok {
				t.Fatalf("%s\n%+v", msg, errResponseor)
			}
			t.Fatal(msg, err)
		}

		log.Println(paths)
	})
}

// TestSites tests GoatCounter Sites endpoints
func TestSites(t *testing.T) {
	code := os.Getenv("CODE")
	token := os.Getenv("TOKEN")
	client := NewClient(code, token)

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

		log.Println(site)
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

		log.Println(sites)
	})
}

// TestStats tests GoatCounter Statistics endpoints
func TestStats(t *testing.T) {
	code := os.Getenv("CODE")
	token := os.Getenv("TOKEN")
	client := NewClient(code, token)

	t.Run("hits", func(t *testing.T) {
		hits, err := client.Stats().Hits()
		if err != nil {
			msg := "expected success"
			if errResponseor, ok := err.(*ErrorResponse); ok {
				t.Fatalf("%s\n%+v", msg, errResponseor)
			}
			t.Fatal(msg, err)
		}

		log.Println(hits)
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

		log.Println(total)
	})
}

// TestUsers tests GoatCounter Users endpoints
func TestUsers(t *testing.T) {
	code := os.Getenv("CODE")
	token := os.Getenv("TOKEN")
	client := NewClient(code, token)

	t.Run("me", func(t *testing.T) {
		user, err := client.Users().Me()
		if err != nil {
			msg := "expected success"
			if errResponseor, ok := err.(*ErrorResponse); ok {
				t.Fatalf("%s\n%+v", msg, errResponseor)
			}
			t.Fatal(msg, err)
		}

		log.Println(user)
	})
}
