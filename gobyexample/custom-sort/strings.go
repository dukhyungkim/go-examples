package main

type Strings []string

func (s Strings) Len() int {
	return len(s)
}

func (s Strings) Less(i, j int) bool {
	return len(s[j]) < len(s[i])
}

func (s Strings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
