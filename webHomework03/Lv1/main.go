package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
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
	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Println(string(all))
	html := string(all)
	title := regexp.MustCompile(`<h2 class="titleview">(?:(.+?))</h2>`)
	matchTitle := title.FindAllStringSubmatch(html, -1)
	var txtTitle string
	for _, v := range matchTitle {
		txtTitle = v[1]
	}
	regexContent := `</p>(?:(.+?))<!--listE-->`
	cRex, err := regexp.Compile(regexContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	content := cRex.FindAllStringSubmatch(html, -1)
	var txtContent string
	for _, v := range content {
		txtContent = v[1]
	}
	result := txtTitle + "\n\r" + strings.ReplaceAll(txtContent, "<br>", "\n")
	result = txtTitle + "\n\r" + strings.ReplaceAll(result, "<br/>", "\n")
	result = txtTitle + "\n\r" + strings.ReplaceAll(result, "<!--listS-->", "\n")
	create, err := os.Create("C:/Users/诚/Desktop/新建 文本文档 (5).txt")
	if err != nil {
		panic(err)
		return
	}
	defer create.Close()
	writeString, err := create.WriteString(result)
	if err != nil {
		return
	}
	fmt.Println(writeString)

	//defer res.Body.Close()
	//if res.StatusCode != 200 {
	//	fmt.Println("status is not ", 200)
	//}
	//
	//doc, err := goquery.NewDocumentFromReader(res.Body)
	//if err != nil {
	//	panic(err)
	//}
	//doc.Find(".content").Each(func(i int, selection *goquery.Selection) {
	//	title := ""
	//	content := ""
	//	title = strings.TrimSpace(selection.Find("h2").Text())
	//	fmt.Println(title)
	//	content = strings.TrimSpace(selection.Text())
	//	create, err := os.Create("C:/Users/诚/Desktop/新建 文本文档 (5).txt")
	//	if err != nil {
	//		panic(err)
	//		return
	//	}
	//	defer create.Close()
	//	writeString, err := create.WriteString(content)
	//	if err != nil {
	//		return
	//	}
	//	fmt.Println(writeString)
	//})
}
