package engine

import "crawler/distributed_crawler/config"

//为worker使用RPC 包装parser
type Parser interface { //序列化
	Parse([]byte, string) ParserRestul
	Serialized() (string, interface{})
}

type Request struct {
	Url   string
	Parse Parser
}
type ParserRestul struct {
	Items    []Item
	Requests []Request
}

type Item struct {
	Url     string
	Id      string
	Payload interface{}
}

type NilFun struct {
}

func (n NilFun) Parse(bytes []byte, s string) ParserRestul {
	return ParserRestul{}
}
func (n NilFun) Serialized() (string, interface{}) {
	return config.NilParser, nil
}

type ParseFun func([]byte, string) ParserRestul

//包装parse
type FuncParse struct {
	parser ParseFun
	name   string
}

func (f *FuncParse) Parse(contents []byte, url string) ParserRestul {
	return f.parser(contents, url)
}

func (f *FuncParse) Serialized() (string, interface{}) {
	return f.name, nil
}

//工厂化
func NewFuncParse(p ParseFun, name string) *FuncParse {
	return &FuncParse{
		parser: p,
		name:   name,
	}
}
