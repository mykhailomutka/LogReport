package report

import (
	"bufio"
	"os"

	"github.com/your-org/logreport/internal/parse"
	"github.com/your-org/logreport/internal/util"
)

type Report struct {
	InputFile string `json:"inputFile"`

	TotalLines    int `json:"totalLines"`
	ParsedLines   int `json:"parsedLines"`
	UnparsedLines int `json:"unparsedLines"`

	LevelCounts  map[string]int `json:"levelCounts"`
	UserCounts   map[string]int `json:"userCounts"`
	ActionCounts map[string]int `json:"actionCounts"`
	UniqueIPs    []string       `json:"uniqueIps"`
	Emails       []string       `json:"emails"`
}

func BuildFromFile(path string, filter Filter, lp parse.LineParser) (Report, error) {
	f, err := os.Open(path)
	if err != nil {
		return Report{}, err
	}
	defer f.Close()

	r := Report{
		InputFile:    path,
		LevelCounts:  make(map[string]int),
		UserCounts:   make(map[string]int),
		ActionCounts: make(map[string]int),
	}

	ipSet := util.NewStringSet()
	emailSet := util.NewStringSet()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		r.TotalLines++

		e := lp.ParseLine(line)
		if e.Level == "" && e.User == "" && e.IP == "" && e.Action == "" && e.Timestamp == nil {
			r.UnparsedLines++
			for _, em := range parse.ExtractEmails(line) {
				emailSet.Add(em)
			}
			continue
		}

		if !passesFilter(e, filter) {
			continue
		}

		r.ParsedLines++
		r.LevelCounts[e.Level]++
		if e.User != "" {
			r.UserCounts[e.User]++
		}
		if e.Action != "" {
			r.ActionCounts[e.Action]++
		}
		if e.IP != "" {
			ipSet.Add(e.IP)
		}
		for _, em := range parse.ExtractEmails(line) {
			emailSet.Add(em)
		}
	}
	if err := sc.Err(); err != nil {
		return Report{}, err
	}

	r.UniqueIPs = ipSet.Sorted()
	r.Emails = emailSet.Sorted()
	return r, nil
}

func passesFilter(e parse.Entry, f Filter) bool {
	if f.Level != "" && e.Level != f.Level {
		return false
	}
	if f.Since != nil && e.Timestamp != nil && e.Timestamp.Before(*f.Since) {
		return false
	}
	if f.Until != nil && e.Timestamp != nil && e.Timestamp.After(*f.Until) {
		return false
	}
	return true
}
