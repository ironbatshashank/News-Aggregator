package main

import ("fmt",
		"net/http")

func main(){
	http.HandlerFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
