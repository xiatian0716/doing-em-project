package IPProxyUtils

import (
	"go-crawler-test/utils"
	"io/ioutil"
	"log"
	"regexp"
	"testing"
)

func TestIndexParse(t *testing.T) {
	// 文件读取
	iot, _ := ioutil.ReadFile("ip.json")
	// 去除空格或制表符
	str := utils.CompressStr(string(iot))
	log.Printf("%s", str)

	// 匹配内容
	// <span>盈余公积</span></td>{{each data as value i}}<td class="tips-data-Right"><span>{{formatRate(value.SURPLUSRESERVE_YOY)}}</span>
	// <spanstyle="display:none;">其中:其他权益工具</span></td>{{eachdataasvaluei}}<tdclass="tips-data-Right"><span>{{formatNumber(value.OTHEREQUITYOTHER,2)}}</span>
	// {"ip":"117.69.129.221","port":4527}
	var intoRe = regexp.MustCompile(`{"ip":"([^"]+)","port":([^}]+)`)

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
		modelGetMap[string(match2[1])] = string(match2[2])
		endPoint := len(modelGetMap)

		// 储存
		if endPoint > startPoint {
			modelGetSli = append(modelGetSli, string(match2[1])+"-"+string(match2[2]))
		}

		t.Logf("%d#%s:%s", i, string(match2[1]), string(match2[2]))
	}
	t.Logf("%s", modelGetSli)
}
