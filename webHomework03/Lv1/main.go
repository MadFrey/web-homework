package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strings"
)

func main() { //爬表单
	funny := "http://xiaodiaodaya.cn/xh/sh/4.html"
	res, err := http.Get(funny)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("status is not ", 200)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	doc.Find(".content").Each(func(i int, selection *goquery.Selection) {
		title := ""
		content := ""
		title = strings.TrimSpace(selection.Find("h2").Text())
		fmt.Println(title)
		content = strings.TrimSpace(selection.Text())
		create, err := os.Create("C:/Users/诚/Desktop/新建 文本文档 (5).txt")
		if err != nil {
			panic(err)
			return
		}
		defer create.Close()
		writeString, err := create.WriteString(content)
		if err != nil {
			return
		}
		fmt.Println(writeString)
	})
}
