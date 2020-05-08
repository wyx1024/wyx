package main

import (
	concurrent_controller "crawler/concurrent_crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("concurrent_crawler/frontend/view/")))
	header, err := concurrent_controller.CreateSearchResultHandler("concurrent_crawler/frontend/view/template.html")
	if err != nil {
		panic(err)
	}
	http.Handle("/car", header)
	http.ListenAndServe(":1024", nil)

}
