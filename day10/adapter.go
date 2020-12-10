package main

import (
	"fmt"
	"../lib"
  "strconv"
  "sort"
)

func main() {
	lines := lib.ReadInput()

	ints := []int{}
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		lib.Check(err)
		ints = append(ints, i)
	}
  fmt.Println(part1(ints))
}

func part1(jolts []int) int{
  sort.Ints(jolts)
  count1, count3 := 1, 1
  fmt.Println(jolts)
  for i := 1; i < len(jolts); i++ {
    if(jolts[i]-jolts[i-1] == 1){
      count1++
    } else if(jolts[i]-jolts[i-1] == 3){
      count3++
    }
  }
  return count1*count3
}
