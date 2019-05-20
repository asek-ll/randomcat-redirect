package main

import (
	"fmt"
	"net/http"
    "encoding/json"
    "io/ioutil"

    "google.golang.org/appengine"
)

type CatImage struct {
	File []string `json:"rile"`
}

func main() {
	http.HandleFunc("/random-cat", handle)
    appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")

	resp, _ := http.Get("http://aws.random.cat/meow")

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var image CatImage

	json.Unmarshal(body, &image)

	fmt.Fprintln(w, image.File)
}
