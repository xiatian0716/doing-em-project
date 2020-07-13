package parseModel

import (
	"github.com/chromedp/chromedp"
)

// 输出
func outerHTML(resq *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.OuterHTML("body > div.hlbody > div.box.all.type1 > div.gbboxb > div.ngbggulbody.list.clearfix > div", resq),
	}
}
