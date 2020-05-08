package concurrent_engine

type Request struct {
	Url      string
	ParseFun ParseFun
}
type ParserRestul struct {
	Items    []Item
	Requests []Request
}

type ParseFun func([]byte, string) ParserRestul

func NilFun([]byte) ParserRestul {
	return ParserRestul{}
}

type Item struct {
	Url     string
	Id      string
	Payload interface{}
}
