package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//fetch, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	fetch, err := ioutil.ReadFile("citylist_test_data.html")
	
	if err != nil{
		panic(err)
	}
	
	//fmt.Printf("%s\n", fetch)
	result := ParseCityList(fetch)
	
	const resultSize = 470
	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	var expectedCitys = []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}
	if len(result.Requests) != resultSize{
		t.Errorf("result shuold have %d request; but had %d\n", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize{
		t.Errorf("result shuold have %d Items; but had %d\n", resultSize, len(result.Items))
	}
	for i, url := range expectedUrls{
		if result.Requests[i].Url != url{
			t.Errorf("expected url #%d; %s; but was %s\n", i, url, result.Requests[i].Url)
		}
	}
	for i, item := range expectedCitys{
		if result.Items[i].(string) != item{
			t.Errorf("expected city #%d; %s; but was %s\n", i, item, result.Items[i].(string))
		}
	}
}