package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item count: #%d: %v", itemCount, item)
			itemCount++

			save(item)
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false)) //must turn off in docker

	if err != nil {
		return "", err
	}

	resp, err := client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	//fmt.Printf("%+v",resp)
	return resp.Id, err
}
