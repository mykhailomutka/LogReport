package parse

import "time"

type Entry struct {
	Raw       string
	Timestamp *time.Time
	Level     string
	User      string
	Action    string
	IP        string
	Message   string
}
