# LogReport (Go)
A production-style Go CLI project that reads a log file, extracts structured fields, and generates a report.
It supports text output and JSON output, plus basic filters.

## Quick start
1. Install Go 1.22+
2. Run a report
   go run ./cmd/logreport -in ./examples/sample.log

## CLI
1. Text report
   go run ./cmd/logreport -in ./examples/sample.log

2. JSON report
   go run ./cmd/logreport -in ./examples/sample.log -json

3. Filter by level
   go run ./cmd/logreport -in ./examples/sample.log -level ERROR

4. Filter by time (RFC3339)
   go run ./cmd/logreport -in ./examples/sample.log -since 2026-03-29T10:22:00Z

## Build
1. go build ./cmd/logreport

## Test
1. go test ./...
