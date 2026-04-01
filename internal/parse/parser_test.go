package parse

import "testing"

func TestParseLine(t *testing.T) {
	p := DefaultLineParser()

	e := p.ParseLine("2026-03-29 10:22:30 ERROR user=alex action=checkout ip=10.0.0.2 msg=\"payment failed\"")
	if e.Level != "ERROR" {
		t.Fatalf("expected ERROR, got %q", e.Level)
	}
	if e.User != "alex" {
		t.Fatalf("expected user alex, got %q", e.User)
	}
	if e.IP != "10.0.0.2" {
		t.Fatalf("expected ip 10.0.0.2, got %q", e.IP)
	}
	if e.Message != "payment failed" {
		t.Fatalf("expected message, got %q", e.Message)
	}
}

func TestExtractEmails(t *testing.T) {
	emails := ExtractEmails("contact a@b.com and c.d+e@x.io now")
	if len(emails) != 2 {
		t.Fatalf("expected 2 emails, got %d", len(emails))
	}
}
