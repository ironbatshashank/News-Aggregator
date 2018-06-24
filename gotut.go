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

func main() {
	a_car := car{gas_pedal: 22341, 
				brake_pedal: 0, 
				steering_wheel: 12561, 
				top_speed_kmh: 225.0}

	//this is same as
	//a_car := car{22341, 0, 1256, 225.0}

	fmt.Println(a_car.gas_pedal)
}