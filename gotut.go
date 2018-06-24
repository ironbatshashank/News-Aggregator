package main

import ("fmt"
		"encoding/xml"
		"io/ioutil"
		"net/http")

type Sitemapindex struct {
	Locations []Location 'xml:"sitemap"'
}

type Location struct {
	Loc string 'xml:"loc"'
}

func(L Location) String() string {
	return fmt.Sprintf(1.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washintonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	var s Sitemapindex
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}
}