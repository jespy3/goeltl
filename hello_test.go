package main

import (
	"fmt"
	"testing"
)

func FailedTest(actual, expected string) string {
  return fmt.Sprintf("Got: %v, Expected: %v", actual, expected)
}

func TestWritesHelloWorld(t *testing.T) {
  ans := WritesHelloWorld()
  if ans != "Hello World!" {
    t.Fatalf("\n" +
      "have WritesHelloWorld() = \"%s\"\n" +
      "want WritesHelloWorld() = \"Hello World!\"", ans)
  }
}

func TestBaseUrl(t *testing.T) {
  t.Run("returns 200", func(t *testing.T) {
    // given
    uri := "http://pokeapi.co/api/v2/pokedex/kanto/"

    // when
    res := getAllPokemon(uri)

    // then
    actual := fmt.Sprintf("%v", res.StatusCode)
    expected := "200"

    if actual != expected {
      t.Errorf(FailedTest(actual, expected))
    }
  })
}

// func TestSpeaker(t *testing.T) {
//   t.Run("calls Speaker", func(t *testing.T) {
//     f := &fakeSpeaker{}

//     Hello(f, "GoCon")

//     actual := f.CallCount
//     expected := 1

//     if actual != expected {
//       t.Errorf("Got: %v, Expected: %v", actual, expected)
//     }
//   })
// }