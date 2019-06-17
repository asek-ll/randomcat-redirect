package main

import (
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"io/ioutil"
	"net/http"
)

type CatImage struct {
	File string `json:"file"`
}

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}
func handle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get("http://aws.random.cat/meow")
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var image CatImage
	json.Unmarshal(body, &image)

	w.Header().Set("Location", image.File)
    w.Header().Set("Cache-Control", "max-age=60, must-revalidate")
	w.WriteHeader(http.StatusFound)
}
