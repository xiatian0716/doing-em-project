package utils

import (
	"go-crawler-test/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var fetchBodyRateLimiter = time.Tick(config.RateLimiter)

func FetchBody(url string) ([]byte, error) {
	<-fetchBodyRateLimiter
	log.Printf("Fetching url:%s", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Http get err:", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Http status code:", resp.StatusCode)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
