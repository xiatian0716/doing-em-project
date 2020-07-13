package parseModel

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"go-crawler-test/utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"testing"
)

func TestGetList(t *testing.T) {
	url := "https://guba.eastmoney.com/remenba.aspx?type=1"

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
	parseItemList(utils.CompressStr(resq))
}

// 获取Index
func TestGetIndex(t *testing.T) {
	index, _ := utils.FetchBody("https://emweb.securities.eastmoney.com/NewFinanceAnalysis/Index?type=web&code=sh605001")
	log.Printf("%s", index)
}

// 解析Index
func TestIndexParse(t *testing.T) {
	// 文件读取
	iot, _ := ioutil.ReadFile("Index.html")
	// 去除空格或制表符
	str := utils.CompressStr(string(iot))
	//log.Printf("%s", str)

	// 匹配内容
	// <span>盈余公积</span></td>{{each data as value i}}<td class="tips-data-Right"><span>{{formatRate(value.SURPLUSRESERVE_YOY)}}</span>
	// <spanstyle="display:none;">其中:其他权益工具</span></td>{{eachdataasvaluei}}<tdclass="tips-data-Right"><span>{{formatNumber(value.OTHEREQUITYOTHER,2)}}</span>
	var intoRe = regexp.MustCompile(`[<span|<spanstyle].*?>([^<]+)</span></td>{{eachdataasvaluei}}.*?value\.([^)|^,]+)`)

	// 完整匹配项
	match := intoRe.FindAll([]byte(str), 900)
	t.Logf("%d", len(match))

	// 子匹配项
	modelGetSli := []string{}
	modelGetMap := make(map[string]string)
	for i := 0; i < len(match); i++ {
		match2 := intoRe.FindSubmatch(match[i])

		// 去重
		startPoint := len(modelGetMap)
		modelGetMap[string(match2[2])] = string(match2[1])
		endPoint := len(modelGetMap)

		// 储存
		if endPoint > startPoint {
			modelGetSli = append(modelGetSli, string(match2[2])+"-"+string(match2[1]))
		}

		t.Logf("%d#%s--%s", i, string(match2[2]), string(match2[1]))
	}
	t.Logf("%s", modelGetSli)
}

func TestHttpproxy(t *testing.T) {
	urli := url.URL{}
	//设置一个http代理服务器格式
	urlproxy, _ := urli.Parse("117.69.97.175:4536")
	//设置一个http客户端
	client := &http.Client{
		Transport: &http.Transport{ //设置代理服务器
			Proxy: http.ProxyURL(urlproxy),
		},
	}
	//访问地址http://myip.top
	rqt, err := http.NewRequest("GET", "http://www.baidu.com/", nil)
	if err != nil {
		println("接口获取IP失败!")
		return
	}
	//添加一个识别信息
	rqt.Header.Add("User-Agent", "Lingjiang")
	//处理返回结果
	response, _ := client.Do(rqt)
	defer response.Body.Close()
	//读取内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	//显示获取到的IP地址
	fmt.Println("http:", string(body))
	return

}
