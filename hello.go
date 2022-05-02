package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
  fmt.Println("Welcome to Goeltl, the Poeltl helper.")
  fmt.Println("Make your first guess of an NBA player...")

  fmt.Println("Jokes. Let's start with Pokemon for now...")
  uri := "http://pokeapi.co/api/v2/pokedex/kanto/"
  pokemon := getAllPokemon(uri)
  fmt.Println(pokemon)
}

func WritesHelloWorld() string {
  result := "Hello World!"
  return result
}

type pokemon struct {
  entry_number    int
  pokemon_species species
}

type species struct {
  name string
  url  string
}

func getAllPokemon(uri string) *http.Response {
  res, err := http.Get(uri)

  if err != nil {
    fmt.Print(err.Error())
  }

  // fmt.Printf("response = %v", res)
  fmt.Printf("res.Status = %v\n", res.Status)
  fmt.Printf("res.StatusCode = %v\n", res.StatusCode)

  data, err := ioutil.ReadAll(res.Body)

  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(data))

  return res
}