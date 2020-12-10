package main

import (
	"fmt"
	"../lib"
  "strconv"
)

func main() {
	lines := lib.ReadInput()

	ints := []int{}
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		lib.Check(err)
		ints = append(ints, i)
	}
  fmt.Println(firstBadNumber(ints))
}

func firstBadNumber(ints []int) int {
  preamble := 25
	for i, _ := range(ints) {
    if(i+preamble == len(ints)){
      break
    }
    if(!containsSum(ints[i:i+preamble], ints[i+preamble])){
      return ints[i+preamble]
      break;
    }
  }
  return 0
}

func containsSum(ints []int, curr int) bool {
	for i, int1 := range ints{
		for _, int2 := range ints[i:]{
			if int1+int2 == curr {
				return true
			}
		}
	}
	return false
}
