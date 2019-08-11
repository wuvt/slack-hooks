package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/wuvt", wuvtHandler)
	http.HandleFunc("/yi", yiHandler)
	http.ListenAndServe(":8080", nil)
}
