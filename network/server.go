package network

import (
	"fmt"
	"net/http"
	"strconv"
)

var htmlString = "<h1>Not found</h1>"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlString)
}

func ServePage(html string, port int) {
	fmt.Println("Starting server")
	htmlString = html
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	fmt.Println(err.Error())

}
