package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strings"
	"strconv"
)

func isValid(input string) (bool) {
  words := SplitAny(input, "- : ")

  first, err := strconv.Atoi(words[0])
  second, err1 := strconv.Atoi(words[1])
	if (err != nil && err1 == nil) {
		log.Fatalf("invalid indexes: %s", err)
	}
  passwordRune := []rune(words[3])

  if(string(passwordRune[first-1])==words[2] && string(passwordRune[second-1])!=words[2] ||
      string(passwordRune[first-1])!=words[2] && string(passwordRune[second-1])==words[2]){
    return true
  }
  return false
}

func SplitAny(s string, seps string) []string {
    splitter := func(r rune) bool {
        return strings.ContainsRune(seps, r)
    }
    return strings.FieldsFunc(s, splitter)
}

func main() {
  //Read the file containing input
  file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  var text []string

  for scanner.Scan() {
      text = append(text, scanner.Text())
  }

  file.Close()
  var count = 0
  for _, each_ln := range text {
    if(isValid(each_ln)){
      count++
    }
  }
  fmt.Println("Total valid passwords are :" , count)
}
