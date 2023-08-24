package goatcounter

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestTypePaths(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error("expected to be able to determine working dir")
	}

	t.Run("list", func(t *testing.T) {
		name := filepath.Join(dir, "../samples/paths.list.json")
		b, err := os.ReadFile(name)
		if err != nil {
			t.Error("expected file to exist")
		}

		pathsResponse := &PathsResponse{}
		if err := json.Unmarshal(b, pathsResponse); err != nil {
			t.Error("expected content to unmarshal as User", err)
		}
	})
}
func TestTypeSites(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error("expected to be able to determine working dir")
	}

	t.Run("get", func(t *testing.T) {
		name := filepath.Join(dir, "../samples/sites.get.json")
		b, err := os.ReadFile(name)
		if err != nil {
			t.Error("expected file to exist")
		}

		site := &Site{}
		if err := json.Unmarshal(b, site); err != nil {
			t.Error("expected content to unmarshal as User", err)
		}
	})
	t.Run("list", func(t *testing.T) {
		name := filepath.Join(dir, "../samples/sites.list.json")
		b, err := os.ReadFile(name)
		if err != nil {
			t.Error("expected file to exist")
		}

		sitesRespone := &SitesResponse{}
		if err := json.Unmarshal(b, sitesRespone); err != nil {
			t.Error("expected content to unmarshal as User", err)
		}
	})
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

		hitsResponse := &HitsResponse{}
		if err := json.Unmarshal(b, hitsResponse); err != nil {
			t.Error("expected content to unmarshal as User", err)
		}
	})
	t.Run("total", func(t *testing.T) {
		name := filepath.Join(dir, "../samples/stats.total.json")
		b, err := os.ReadFile(name)
		if err != nil {
			t.Error("expected file to exist")
		}

		total := &Total{}
		if err := json.Unmarshal(b, total); err != nil {
			t.Error("expected content to unmarshal as User", err)
		}
	})
}
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
