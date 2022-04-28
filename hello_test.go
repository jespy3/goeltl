package main

import (
  "testing"
)

func TestWritesHelloWorld(t *testing.T) {
  ans := WritesHelloWorld()
  if ans != "Hello World!" {
    t.Fatalf("\n" +
      "have WritesHelloWorld() = \"%s\"\n" +
      "want WritesHelloWorld() = \"Hello World!\"", ans)
  }
}
