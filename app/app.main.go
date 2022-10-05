package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
	"strings"
)

func main() {
	var url string
	fmt.Println("Введите url поста")
	fmt.Scanf("%s\n", &url)

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
		ParseFunc: parseHabrPost,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}

func parseHabrPost(g *geziyor.Geziyor, r *client.Response) {
	var topic []string
	r.HTMLDoc.Find("div.article-formatted-body div").Each(func(i int, s *goquery.Selection) {

		if s.Find("h3") != nil {
			s.Find("h3").Each(func(i int, s *goquery.Selection) {
				topic = append(topic, s.Text())
			})
		}

		if s.Find("p strong") != nil {
			s.Find("p strong").Each(func(i int, s *goquery.Selection) {
				topic = append(topic, s.Text())
			})
		}

	})

	g.Exports <- map[string]interface{}{
		"author": strings.Trim(r.HTMLDoc.Find("a.tm-user-info__username").Text(), "\n "),
		"title":  r.HTMLDoc.Find("h1.tm-article-snippet__title_h1 span").Text(),
		"body":   r.HTMLDoc.Find("div.article-formatted-body div").Text(),
		"topic":  topic,
	}
}
