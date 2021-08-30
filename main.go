package main

import (
	"net/http"
)

//var temp = template.Must()
func main() {
	http.ListenAndServe(":8080", nil)
}
