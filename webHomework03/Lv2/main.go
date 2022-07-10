package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	for i := 1; i < 6; i++ {
		url := "https://fabiaoqing.com/biaoqing/lists/page/" + strconv.Itoa(i) + ".html"
		client := &http.Client{
			Timeout: 2 * time.Second,
		}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36")
		req.Header.Set("Host", "fabiaoqing.com")
		res, err := client.Do(req)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			panic(err)
			return
		}
		doc.Find(".tagbqppdiv").Each(func(i int, selection *goquery.Selection) {
			src, _ := selection.Find("img").Attr("data-original")
			res, err := http.Get(src)
			if err != nil {
				return
			}
			all, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
				return
			}
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			name := fmt.Sprintf("%06v", rnd.Int31n(1000000))
			err = ioutil.WriteFile(fmt.Sprintf("C:/Users/诚/Desktop/picture/%s 1_%d.jpg", name, i), all, 0644)//注意一下路径
			if err != nil {
				panic(err)
				return
			}
			fmt.Println(src)
		})
	}
}

//.Find(".container").Find(".ui right floated grid").Find("right floated left aligned sixteen wide column hotbq").Find(".ui segment imghover").
//Find(".tagbqppdiv").Find("a").
