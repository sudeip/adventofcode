package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "io"
	"strconv"
)

func ReadInts(r io.Reader) ([]int, error) {
  scanner := bufio.NewScanner(r)
  scanner.Split(bufio.ScanWords)
  var result []int
  for scanner.Scan() {
      x, err := strconv.Atoi(scanner.Text())
      if err != nil {
          return result, err
      }
      result = append(result, x)
  }
  return result, scanner.Err()
}

func main() {
  //Making this a simple one file solution to print part1 and part 2 answers
  //Read the file containing input
  file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
  //store the data in integer array
  ints, err := ReadInts(file)

  //nested for loops to determine numbers that add upto 2020 and then print
  //the data file given would work fine however imagine 1010 or 505 being in the input list

  // part1:
  // for _, first := range ints {
  //   for _, second := range ints {
  //     if (first+second == 2020){
	// 	    fmt.Println("part1 match numbers  ", first, second)
	// 	    fmt.Println("part1 answer         ", first*second)
  //       break part1;
  //     }
  //   }
  // }
  //
  // part2:
  // for _, first := range ints {
  //   for _, second := range ints {
  //     for _, third := range ints {
  //       if (first+second+third == 2020){
  // 		    fmt.Println("part2 match numbers  ", first, second, third)
  // 		    fmt.Println("part2 answer         ", first*second*third)
  //         break part2;
  //       }
  //     }
  //   }
  // }

  //adding additional check to ensure same number is not counted twice
  var arrLen = len(ints);

  partI:
  for i:=0; i<arrLen; i++{
    for j:=0; j<arrLen; j++{
      if((i != j) && (ints[i]+ints[j] == 2020)){
        fmt.Println("partI match numbers  ", ints[i], ints[j])
        fmt.Println("partI answer         ", ints[i]*ints[j])
        break partI;
      }
    }
  }

  partII:
  for i:=0; i<arrLen; i++{
    for j:=0; j<arrLen; j++{
      for k:=0; k<arrLen; k++{
        if((i != j) && (j != k) && (k != i) && (ints[i]+ints[j]+ints[k] == 2020)){
          fmt.Println("partII match numbers  ", ints[i], ints[j], ints[k])
          fmt.Println("partII answer         ", ints[i]*ints[j]*ints[k])
          break partII;
        }
      }
    }
  }
}
