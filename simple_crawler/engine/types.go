package simple_engine

type Request struct {
	Url      string
	ParseFun ParseFun
}
type ParserRestul struct {
	Items    []interface{}
	Requests []Request
}

type ParseFun func([]byte) ParserRestul

func NilFun([]byte) ParserRestul {
	return ParserRestul{}
}
