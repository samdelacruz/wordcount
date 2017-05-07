package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/samdelacruz/wordcount"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	counts := wordcount.FromString(string(body))

	jsonStr, err := json.Marshal(counts)
	fmt.Fprintf(w, "%s", jsonStr)
}

func main() {
	//text := "The quick brown fox jumped over the lazy dog"

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
