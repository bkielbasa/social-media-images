package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/imglabel", imgLabel)
	http.ListenAndServe(":8090", nil)
}
