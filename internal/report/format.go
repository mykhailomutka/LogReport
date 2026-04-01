package report

import (
	"fmt"
	"sort"
	"strings"
)

type kv struct {
	k string
	v int
}

func FormatText(r Report) string {
	var b strings.Builder

	fmt.Fprintf(&b, "==== LogReport ====%s", "\n")
	fmt.Fprintf(&b, "Input: %s%s", r.InputFile, "\n")
	fmt.Fprintf(&b, "Total lines: %d%s", r.TotalLines, "\n")
	fmt.Fprintf(&b, "Parsed lines: %d%s", r.ParsedLines, "\n")
	fmt.Fprintf(&b, "Unparsed lines: %d%s", r.UnparsedLines, "\n")
	fmt.Fprintf(&b, "%s", "\n")

	fmt.Fprintf(&b, "Levels:%s", "\n")
	for _, k := range []string{"INFO", "WARN", "ERROR"} {
		if v, ok := r.LevelCounts[k]; ok {
			fmt.Fprintf(&b, "  %s: %d%s", k, v, "\n")
		}
	}
	fmt.Fprintf(&b, "%s", "\n")

	fmt.Fprintf(&b, "Top Users:%s", "\n")
	for _, kv := range topN(r.UserCounts, 10) {
		fmt.Fprintf(&b, "  %s: %d%s", kv.k, kv.v, "\n")
	}
	fmt.Fprintf(&b, "%s", "\n")

	fmt.Fprintf(&b, "Top Actions:%s", "\n")
	for _, kv := range topN(r.ActionCounts, 10) {
		fmt.Fprintf(&b, "  %s: %d%s", kv.k, kv.v, "\n")
	}
	fmt.Fprintf(&b, "%s", "\n")

	fmt.Fprintf(&b, "Unique IPs (%d):%s", len(r.UniqueIPs), "\n")
	for _, ip := range r.UniqueIPs {
		fmt.Fprintf(&b, "  %s%s", ip, "\n")
	}
	fmt.Fprintf(&b, "%s", "\n")

	fmt.Fprintf(&b, "Emails (%d):%s", len(r.Emails), "\n")
	for _, e := range r.Emails {
		fmt.Fprintf(&b, "  %s%s", e, "\n")
	}

	return b.String()
}

func topN(m map[string]int, n int) []kv {
	out := make([]kv, 0, len(m))
	for k, v := range m {
		out = append(out, kv{k: k, v: v})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].v == out[j].v {
			return out[i].k < out[j].k
		}
		return out[i].v > out[j].v
	})
	if n > 0 && len(out) > n {
		out = out[:n]
	}
	return out
}
