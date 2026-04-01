package util

import "sort"

type StringSet struct {
	m map[string]struct{}
}

func NewStringSet() *StringSet {
	return &StringSet{m: make(map[string]struct{})}
}

func (s *StringSet) Add(v string) {
	if v == "" {
		return
	}
	s.m[v] = struct{}{}
}

func (s *StringSet) Sorted() []string {
	out := make([]string, 0, len(s.m))
	for k := range s.m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
