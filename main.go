package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/blog-post", blogPost)
	http.ListenAndServe(fmt.Sprintf(":%s", env("PORT", "8090")), nil)
}

func env(name, def string) string {
	val := os.Getenv(name)
	if val == "" {
		return def
	}

	return val
}
