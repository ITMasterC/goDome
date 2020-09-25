package persist

import (
	"context"
	"demo-project/crawler/model"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSaver(t *testing.T) {
	profile := model.Profile{
		Age:        34,
		Name:       "天才落寞",
		Gender:     "男",
		Height:     172,
		Income:     "10000-120000",
		Marriage:   "未婚",
		Education:  "本科",
		Occupation: "程序员",
		HoKou:      "广西",
		XinZou:     "水平",
		Blood:      "B",
		XianJu:     "广州",
	}
	id, err := save(profile)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual model.Profile
	err = json.Unmarshal([]byte(resp.Source), &actual)
	if err != nil {
		panic(err)
	}

	if actual != profile {
		t.Errorf("got %v; expected %v", actual, profile)
	}
}
