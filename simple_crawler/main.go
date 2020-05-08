package main

import (
	simple_config "crawler/simple_crawler/config"
	simple_engine "crawler/simple_crawler/engine"
	simple_parser "crawler/simple_crawler/parser"
)

func main() {
	r := simple_engine.Request{
		Url:      simple_config.URL,
		ParseFun: simple_parser.ParserCarTypeList,
	}
	simple_engine.Run(r)
}
