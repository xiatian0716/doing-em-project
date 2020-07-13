package parse

import (
	"bytes"
	"context"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/chromedp/chromedp"
	"go-crawler-test/utils"
	"log"
	"strings"
	"testing"
	"time"
)

func TestXPath(t *testing.T) {
	body, _ := utils.FetchBody("https://guba.eastmoney.com/remenba.aspx?type=1")
	doc, _ := htmlquery.Parse(bytes.NewReader(body))
	list, _ := htmlquery.QueryAll(doc, `/html/body/div[1]/div[5]/div[2]/div[1]/div/ul/li`)

	for i, n := range list {
		a := htmlquery.FindOne(n, "//a")
		temp := strings.Split(htmlquery.SelectAttr(a, "href"), ",")[1]
		num := strings.Split(temp, ".")[0]
		fmt.Printf("%d %s-----%s\n", i, htmlquery.InnerText(a), num)
	}
}

func TestJson(t *testing.T) {
	body, _ := utils.FetchBody("http://f10.eastmoney.com/BusinessAnalysis/BusinessAnalysisAjax?code=SH600039")
	//log.Println(json.Unmarshal(body, nil))
	log.Println(string(body))
}

func TestChromedp(t *testing.T) {
	var ua string

	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()
	// create context
	ctx, cancel := chromedp.NewContext(c)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.whatsmyua.info/?a`),
		chromedp.WaitVisible(`#custom-ua-string`),
		chromedp.Text(`#custom-ua-string`, &ua),
		chromedp.Sleep(10*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user agent: %s", ua)
}

func DoCrawler() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Sleep(10 * time.Second),
		chromedp.Click(`#zcfzb_next`, chromedp.NodeVisible),
		chromedp.Sleep(3 * time.Second),
		chromedp.Click(`#zcfzb_next`, chromedp.NodeVisible),
		chromedp.Sleep(3 * time.Second),
		chromedp.Click(`#zcfzb_next`, chromedp.NodeVisible),
		chromedp.Sleep(300000 * time.Second),
	}
}
