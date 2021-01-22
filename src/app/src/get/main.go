package main

import (
	"app/src/github.com/PuerkitoBio/goquery"
	"fmt"
)

func main() {
	url := "https://www.hinatazaka46.com/s/official/diary/member/list?ima=0000&ct=8"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find(".p-blog-article").First().Find("img").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("src")
		fmt.Println(url)
	})

	// First()なしで全部取得
	// doc.Find(".p-blog-article").Find("img").Each(func(_ int, s *goquery.Selection) {
	// 	url, _ := s.Attr("src")
	// 	fmt.Println(url)
	// })
}
