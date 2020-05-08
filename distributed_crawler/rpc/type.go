package worker

import (
	"crawler/distributed_crawler/config"
	"crawler/distributed_crawler/engine"
	"crawler/distributed_crawler/parser"
	"errors"
)

type Serialized struct {
	Name string
	Args interface{}
}

type Request struct {
	Url        string
	Serialized Serialized
}
type ParseResult struct {
	Request []Request
	Item    []engine.Item
}

func SerializedRequest(r engine.Request) Request {
	name, args := r.Parse.Serialized()
	return Request{
		Url: r.Url,
		Serialized: Serialized{
			Name: name,
			Args: args,
		},
	}
}
func SerizlizedParseResult(resp engine.ParserRestul) ParseResult {
	result := ParseResult{}
	result.Item = resp.Items
	for _, r := range resp.Requests {
		result.Request = append(result.Request, SerializedRequest(r))
	}
	return result
}

func DeSerializedRequest(r Request) (engine.Request, error) {
	parser, err := deserizliezedParse(r.Serialized)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:   r.Url,
		Parse: parser,
	}, nil
}
func DeSerizlizedParseResult(resp ParseResult) engine.ParserRestul {
	result := engine.ParserRestul{}
	result.Items = resp.Item
	for _, r := range resp.Request {
		deresult, err := DeSerializedRequest(r)
		if err != nil {
			continue
		}
		result.Requests = append(result.Requests, deresult)
	}
	return result
}

func deserizliezedParse(r Serialized) (engine.Parser, error) {

	switch r.Name {
	case config.ParseCarTypeList:
		return engine.NewFuncParse(parser.ParserCarTypeList, r.Name), nil
	case config.ParseCar:
		return engine.NewFuncParse(parser.ParserCar, r.Name), nil
	case config.ParserDetail:
		if args, ok := r.Args.(string); ok {
			return parser.NewProfile(args), nil
		} else {
			return engine.NilFun{}, errors.New("args error")
		}
	case config.NilParser:
		return engine.NilFun{}, nil
	default:
		return engine.NilFun{}, errors.New("unknown parser")
	}

}
