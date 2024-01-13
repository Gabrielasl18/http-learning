package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	UserID int32  `json:"userId"`
	ID     int32  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Fatalln(err)
	}
	for _, post := range posts {
		log.Printf("%+v\n", post)
	}
}
