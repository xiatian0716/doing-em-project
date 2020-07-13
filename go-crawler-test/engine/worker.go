package engine

// 包工头
func Worker(r Request) (ParseResult, error) {
	// Parse网页
	return r.ParseFunc(r.Url), nil
}
