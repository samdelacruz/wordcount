package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/samdelacruz/wordcount/counter"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	bodyStr := string(body)
	counts := counter.FromString(&bodyStr)

	jsonStr, err := json.Marshal(counts)
	fmt.Fprintf(w, "%s", jsonStr)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
