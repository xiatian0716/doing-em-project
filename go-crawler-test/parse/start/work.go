package start

import (
	"context"
	"github.com/chromedp/chromedp"
	"go-crawler-test/engine"
	"go-crawler-test/utils"
	"log"
)

func Woke(url string) engine.ParseResult {
	//url = "https://guba.eastmoney.com/remenba.aspx?type=1"

	result := engine.ParseResult{}

	ctx := context.Background()    // load options
	options := utils.OptionsUtil() // sets options
	ctxOpt, cancelOpt := chromedp.NewExecAllocator(ctx, options...)
	defer cancelOpt()
	ctx, cancel := chromedp.NewContext(ctxOpt) // create context
	defer cancel()

	err := chromedp.Run(ctx, utils.HeadersUtil(
		url, false,
		"guba.eastmoney.com",
		map[string]interface{}{
			"X-Header": "my request header"},
		"cookie1", "value1"))
	if err != nil {
		log.Fatal(err)
	}

	var resq string
	outerHTMLErr := chromedp.Run(ctx, outerHTML(&resq))
	if err != nil {
		log.Fatal(outerHTMLErr)
	}
	parseItemList(utils.CompressStr(resq), &result)

	return result
}
