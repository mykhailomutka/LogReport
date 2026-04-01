package parse

import (
	"regexp"
	"strings"
	"time"
)

type LineParser interface {
	ParseLine(line string) Entry
}

type defaultParser struct {
	reMain *regexp.Regexp
	reMsg  *regexp.Regexp
}

func DefaultLineParser() LineParser {
	return &defaultParser{
		reMain: regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})\s+(\d{2}:\d{2}:\d{2})\s+(INFO|WARN|ERROR)\s+(.*)$`),
		reMsg:  regexp.MustCompile(`msg="([^"]*)"`),
	}
}

func (p *defaultParser) ParseLine(line string) Entry {
	e := Entry{Raw: line}

	m := p.reMain.FindStringSubmatch(line)
	if len(m) == 0 {
		return e
	}

	ts, err := time.Parse("2006-01-02 15:04:05", m[1]+" "+m[2])
	if err == nil {
		e.Timestamp = &ts
	}

	e.Level = m[3]
	rest := m[4]

	if mm := p.reMsg.FindStringSubmatch(rest); len(mm) == 2 {
		e.Message = mm[1]
		rest = p.reMsg.ReplaceAllString(rest, "")
	}

	fields := strings.Fields(rest)
	for _, f := range fields {
		if strings.HasPrefix(f, "user=") {
			e.User = strings.TrimPrefix(f, "user=")
		} else if strings.HasPrefix(f, "action=") {
			e.Action = strings.TrimPrefix(f, "action=")
		} else if strings.HasPrefix(f, "ip=") {
			e.IP = strings.TrimPrefix(f, "ip=")
		}
	}

	if e.Message == "" {
		e.Message = strings.TrimSpace(rest)
	}

	return e
}

func ExtractEmails(text string) []string {
	re := regexp.MustCompile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}`)
	return re.FindAllString(text, -1)
}
