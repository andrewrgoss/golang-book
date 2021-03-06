package main

import (
  "fmt"
  "strings"
)

func main() {
  
  arr := []byte("test") // convert a string to a slice of bytes 
  // [116 101 115 116]
  str := string([]byte{'t','e','s','t'}) // convert a slice of bytes to a string
  // test
    
  fmt.Println(
    // true
    strings.Contains("test", "es"),

    // 2
    strings.Count("test", "t"),

    // true
    strings.HasPrefix("test", "te"),

    // true
    strings.HasSuffix("test", "st"),

    // 1
    strings.Index("test", "e"),

    // "a-b"
    strings.Join([]string{"a","b"}, "-"),

    // == "aaaaa"
    strings.Repeat("a", 5),

    // "bbaa"
    strings.Replace("aaaa", "a", "b", 2),

    // []string{"a","b","c","d","e"}
    strings.Split("a-b-c-d-e", "-"),

    // "test"
    strings.ToLower("TEST"),

    // "TEST"
    strings.ToUpper("test"),
    
    arr, 
    str, 
  )
}