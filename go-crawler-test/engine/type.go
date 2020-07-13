package engine

import (
	"go-crawler-test/model"
)

type Request struct {
	Url       string
	ParseWay  string
	ParseFunc func(Url string) ParseResult
}

type ParseResult struct {
	Requesrts []Request
	Items     []model.Item
}

func NilParse(Url string) ParseResult {
	return ParseResult{}
}
