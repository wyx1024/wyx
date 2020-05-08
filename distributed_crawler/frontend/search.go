package main

import (
	controller "crawler/distributed_crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("distributed_crawler/frontend/view/")))
	header, err := controller.CreateSearchResultHandler("distributed_crawler/frontend/view/template.html")
	if err != nil {
		panic(err)
	}
	http.Handle("/car", header)
	http.ListenAndServe(":1024", nil)

}
