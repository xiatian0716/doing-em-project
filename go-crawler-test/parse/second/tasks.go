package second

import (
	"github.com/chromedp/chromedp"
	"time"
)

// 动作
func actionTasks(resq *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Sleep(3 * time.Second),
		chromedp.Click(`#zcfzb_next`, chromedp.NodeVisible),
		chromedp.OuterHTML("html", resq),
	}
}
