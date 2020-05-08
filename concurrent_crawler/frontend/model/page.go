package concurrent_model

type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []interface{}
}
