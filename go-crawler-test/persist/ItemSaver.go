package persist

import (
	"context"
	"go-crawler-test/config"
	"go-crawler-test/model"
	"log"

	"github.com/olivere/elastic"
)

func ItemSaver() (chan model.Item, error) {
	configs := config.ConfigsParse("./config/config.yaml")
	client, err := elastic.NewClient(
		elastic.SetURL(configs.ESConfig.SetURL),
		// Must turn off sniff in docker
		elastic.SetSniff(configs.ESConfig.SetSniff),
	)
	// 连不上返回err
	if err != nil {
		return nil, err
	}

	outSaveChan := make(chan model.Item)
	go func() {
		itemCount := 0
		for {
			item := <-outSaveChan
			log.Printf("Got item->Item Saver"+"#%d: %s", itemCount, item)
			itemCount++

			// Elasticsearch Clients
			err := save(client, configs, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v:%v", item, err)
			}
		}
	}()
	return outSaveChan, nil
}

func save(client *elastic.Client, configs config.Configs, item model.Item) error {
	indexService := client.Index().
		Index(configs.ESConfig.Index). // database
		Type(configs.ESConfig.Type).   // table
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id) // id
	}
	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
