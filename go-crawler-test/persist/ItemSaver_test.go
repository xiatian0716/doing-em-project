package persist

import (
	"go-crawler-test/config"
	"go-crawler-test/model"
	"go-crawler-test/parse/start"
	"testing"

	"github.com/olivere/elastic"
)

func TestSaver(t *testing.T) {
	configs := config.ConfigsParse("../config/config.yaml")
	expected := model.Item{
		Id: "001",

		Payload: start.Start{
			StartUrl: "https://book.douban.com/tag/诗歌"},
	}

	// 然后我们去拿出来
	// Fetch saved item
	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetURL(configs.ESConfig.SetURL),
		elastic.SetSniff(configs.ESConfig.SetSniff),
	)
	if err != nil {
		panic(err)
	}

	// 测试save方法
	// Save expected item
	err = save(client, configs, expected)
	if err != nil {
		panic(err)
	}
}
