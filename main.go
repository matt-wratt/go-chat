package main

import "net/http"

var api = NewAPI()

func main() {
	http.HandleFunc("/", api.APIHandleFunc)
	http.ListenAndServe(":8082", nil)
}
