package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://www.google.com",
		"http://34.80.248.226/login",
		"http://www.yahoo.com",
		"http://www.facebook.com",
		"http://www.golang.org",
		"http://www.xyz.com",
		"http://www.youtube.com",
		"http://www.facebook.com",
		"http://www.baidu.com",
		"http://www.yahoo.com",
		"http://www.amazon.com",
		"http://www.wikipedia.org",
		"http://www.qq.com",
		"http://www.google.co.in",
		"http://www.twitter.com",
		"http://www.live.com",
		"http://www.taobao.com",
		"http://www.bing.com",
		"http://www.instagram.com",
		"http://www.weibo.com",
		"http://www.sina.com.cn",
		"http://www.linkedin.com",
	}
	urlch := make(chan string)
	respch := make(chan string)

	for i := 0; i < len(urls); i++ {
		go worker(urlch, respch, i)
	}

	for _, url := range urls {
		go generator(url, urlch)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s", <-respch)
	}
}

func generator(url string, urlch chan string) {
	urlch <- url
}

func worker(urlch chan string, resp chan string, id int) {
	url := <-urlch
	res, err := check(url)
	if err == nil {
		resp <- fmt.Sprintf("%s >>>>> workerID:%d >>>>> status: %d\n", url, id, res)
		return
	}
	resp <- fmt.Sprintf("%s  >>>>> in workerID:%d: >>>>> error: %s\n", url, id, err)
}

func check(url string) (status int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}
