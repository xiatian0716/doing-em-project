package parseModel

import (
	"log"
	"regexp"
	"strings"
)

func parseItemList(body string) {
	// 匹配内容
	// <ahref="list,600000.html">(600000)浦发银行</a>
	var intoRe = regexp.MustCompile(`<ahref.*?>([^<]+)`)

	// 完整匹配项
	match := intoRe.FindAll([]byte(body), 9000)
	log.Printf("%d", len(match))

	// 子匹配项
	modelGetSli := []string{}
	modelGetMap := make(map[string]string)
	for i := 0; i < len(match); i++ {
		match2 := intoRe.FindSubmatch(match[i])

		// 去重
		startPoint := len(modelGetMap)
		temp := strings.Split(string(match2[1]), ")")[0]
		if len(strings.Split(temp, "(")) < 2 {
			continue
		}
		code := strings.Split(temp, "(")[1]
		modelGetMap[string(match2[1])] = code
		endPoint := len(modelGetMap)

		// 储存
		if endPoint > startPoint {
			modelGetSli = append(modelGetSli, string(match2[1])+"-"+code)
		}

		log.Printf("%d#%s--%s", i, string(match2[1]), code)
	}
	log.Printf("%s", modelGetSli)
}
