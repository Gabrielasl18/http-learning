package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Pokemon struct {
	name string `json:"name"`
}

func main() {
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return
	}
	fmt.Printf("%s\n", body)

	fmt.Printf("Status: %s\n", res.Status)
}
