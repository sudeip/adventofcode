package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strings"
	"math"
)

func findTrees(input []string, right int, down int)(int){
	  var count = 0
	  var index = right
	  var repeatStr = 1

	  var x = math.Floor(float64(len(input[0])/index))

	  for i:=0+down; i<len(input); {
	    lineStr := strings.Repeat(input[i], repeatStr)
	    lineRune := []rune(lineStr)
	    if(string(lineRune[index]) == "#"){
	      count++
	    }
	    //fmt.Println(i, x, repeatStr, index, count)
	    if(i+1 >= int(x)){
	      repeatStr++
	      x = math.Floor(float64(len(lineStr)/index))
	    }
	    index = index+right
			i = i+down
	  }
		return count
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
	x11 := findTrees(text, 1, 1)
	x31 := findTrees(text, 3, 1)
	x51 := findTrees(text, 5, 1)
	x71 := findTrees(text, 7, 1)
	x12 := findTrees(text, 1, 2)

  fmt.Println("Total trees encountered for 11, 31, 51, 71 and 12 :" , x11, x31, x51, x71, x12)
	fmt.Println("answer for part II : ", x11*x31*x51*x71*x12)
}
