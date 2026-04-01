package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/your-org/logreport/internal/parse"
	"github.com/your-org/logreport/internal/report"
)

type options struct {
	inputPath string
	jsonOut   bool
	level     string
	since     string
	until     string
}

func main() {
	opts := parseFlags()

	if opts.inputPath == "" {
		fmt.Fprintln(os.Stderr, "error: missing input file")
		fmt.Fprintln(os.Stderr, "usage: logreport -in <path> [--json] [--level INFO|WARN|ERROR] [--since RFC3339] [--until RFC3339]")
		os.Exit(2)
	}

	var sincePtr *time.Time
	var untilPtr *time.Time

	if opts.since != "" {
		t, err := time.Parse(time.RFC3339, opts.since)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: invalid --since, expected RFC3339:", err)
			os.Exit(2)
		}
		sincePtr = &t
	}

	if opts.until != "" {
		t, err := time.Parse(time.RFC3339, opts.until)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: invalid --until, expected RFC3339:", err)
			os.Exit(2)
		}
		untilPtr = &t
	}

	filter := report.Filter{
		Level: opts.level,
		Since: sincePtr,
		Until: untilPtr,
	}

	r, err := report.BuildFromFile(opts.inputPath, filter, parse.DefaultLineParser())
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	if opts.jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(r)
		return
	}

	fmt.Print(report.FormatText(r))
}

func parseFlags() options {
	var o options
	flag.StringVar(&o.inputPath, "in", "", "input log file path")
	flag.BoolVar(&o.jsonOut, "json", false, "output JSON instead of text")
	flag.StringVar(&o.level, "level", "", "filter by level (INFO, WARN, ERROR)")
	flag.StringVar(&o.since, "since", "", "filter entries since RFC3339 time")
	flag.StringVar(&o.until, "until", "", "filter entries until RFC3339 time")
	flag.Parse()
	return o
}
