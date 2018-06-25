package main

import ("fmt"
		"net/http"
		"html/template")

type NewsAggPAge struct {
	Title string
	News string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("basictemplating.html")
	fmt.Println(t.Execute(w, p))
}

func main() {
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)

		fmt.Println("\n\n\n",idx)
		fmt.Println("\n",data.Keyword)

		bytes, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("\n\n\n",idx)
		fmt.Println("\n",data.Keyword)
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Titles{
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	for idx, data := range news_map {
		fmt.Println("\n\n\n",idx)
		fmt.Println("\n",data.Keyword)
		fmt.Println("\n",data.Location)
	}
}