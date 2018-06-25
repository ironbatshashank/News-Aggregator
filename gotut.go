package main

import ("fmt"
		"encoding/xml"
		"io/ioutil"
		"net/http")

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}
}