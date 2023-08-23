package goatcounter

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestTypeUser(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error("expected to be able to determine working dir")
	}
	name := filepath.Join(dir, "../samples/me.json")
	b, err := os.ReadFile(name)
	if err != nil {
		t.Error("expected file to exist")
	}

	user := &User{}
	if err := json.Unmarshal(b, user); err != nil {
		t.Error("expected content to unmarshal as User", err)
	}
}
func TestTypeStats(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error("expected to be able to determine working dir")
	}

	t.Run("hits", func(t *testing.T) {
		name := filepath.Join(dir, "../samples/stats.hits.json")
		b, err := os.ReadFile(name)
		if err != nil {
			t.Error("expected file to exist")
		}

		hits := &StatsHits{}
		if err := json.Unmarshal(b, hits); err != nil {
			t.Error("expected content to unmarshal as User", err)
		}
	})
	t.Run("total", func(t *testing.T) {
		name := filepath.Join(dir, "../samples/stats.total.json")
		b, err := os.ReadFile(name)
		if err != nil {
			t.Error("expected file to exist")
		}

		total := &StatsTotal{}
		if err := json.Unmarshal(b, total); err != nil {
			t.Error("expected content to unmarshal as User", err)
		}
	})
}
