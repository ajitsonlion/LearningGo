package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type SiteMap struct {
	Urls []string `xml:"url>loc"`
}

type News struct {
	Titles    []string `xml:"url>news:news>news:title"`
	Keywords  []string `xml:"url>news:news>news:keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Url     string
}

func main() {

	resp, _ := http.Get("https://www.washingtonpost.com/news-blogs-technology-sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	var s SiteMap
	var n News
	newsMap := make(map[string]NewsMap)

	xml.Unmarshal(bytes, &s)

	for _, url := range s.Urls {
		fmt.Println(url)

		urlRes, error := http.Get(strings.TrimSpace(url))
		if error != nil {
			fmt.Println(error)
		}
		urlBytes, _ := ioutil.ReadAll(urlRes.Body)
		xml.Unmarshal(urlBytes, &n)
		fmt.Println(n.Titles)

		for index := range n.Titles {
			fmt.Println(n.Titles)
			newsMap[n.Titles[index]] = NewsMap{n.Keywords[index], n.Locations[index]}

		}

	}

	for index, data := range newsMap {
		fmt.Println("Index", index)
		fmt.Println("Data", data)

		fmt.Println("Locations", data.Keyword)

	}

}
