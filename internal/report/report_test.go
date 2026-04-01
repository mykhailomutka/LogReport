package report

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/your-org/logreport/internal/parse"
)

func TestBuildReport(t *testing.T) {
	tmp := t.TempDir()
	p := filepath.Join(tmp, "a.log")
	err := os.WriteFile(p, []byte("2026-03-29 10:21:00 INFO  user=misha action=login ip=10.0.0.1\n"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	r, err := BuildFromFile(p, Filter{}, parse.DefaultLineParser())
	if err != nil {
		t.Fatal(err)
	}
	if r.TotalLines != 1 {
		t.Fatalf("expected total 1, got %d", r.TotalLines)
	}
	if r.LevelCounts["INFO"] != 1 {
		t.Fatalf("expected INFO=1, got %d", r.LevelCounts["INFO"])
	}
	if r.UserCounts["misha"] != 1 {
		t.Fatalf("expected user misha=1, got %d", r.UserCounts["misha"])
	}
}
