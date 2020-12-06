package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strings"
	"math"
)

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
  var index = 3
  var right = 1

  var x = math.Floor(float64(len(text[0])/index))

  for i:=1; i<len(text); i++{
    lineStr := strings.Repeat(text[i], right)
      fmt.Println(lineStr)
    lineRune := []rune(lineStr)
    if(string(lineRune[index]) == "#"){
      count++
    }
      //fmt.Println(i, x, right, index)
    if(i+1 > int(x)){
      right++
      x = math.Floor(float64(len(lineStr)/index))
    }
    index = index+3
  }

  fmt.Println("Total trees encountered :" , count)
}
