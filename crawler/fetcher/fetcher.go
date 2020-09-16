package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	//resp, err := http.Get("http://www.zhenai.com/zhenghun")
	//<- rateLimiter
	resp, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}
func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, e := reader.Peek(1024)
	if e != nil{
		log.Printf("Fetcher error: %v", e)
		return unicode.UTF8
	}
	e2, _, _ := charset.DetermineEncoding(bytes, "")
	return e2
}
