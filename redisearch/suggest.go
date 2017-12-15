package redisearch

import "sort"

// Suggestion is a single suggestion being added or received from the Autocompleter
type Suggestion struct {
	Term    string  `json:"term"`
	Score   float64 `json:"score"`
	Payload string  `json:"payload"`
}

func (s Suggestion)String() string{
	return SprintInterface(s)
}

// SuggestionList is a sortable list of suggestions returned from an engine
type SuggestionList []Suggestion

func (l SuggestionList) Len() int           { return len(l) }
func (l SuggestionList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l SuggestionList) Less(i, j int) bool { return l[i].Score > l[j].Score } //reverse sorting

// Sort the SuggestionList
func (l SuggestionList) Sort() {
	sort.Sort(l)
}
