package main

import (
	"fmt"
	"strings"
	"../lib"
  "strconv"
)

func main() {
	lines := lib.ReadInput()
  fmt.Println(lastAccumulatorValPart1(lines))

}

func lastAccumulatorValPart1(lines []string) int {
  accumulator := 0
	m := map[int]int{}
	for i:=0;i<len(lines);i++ {
    v := strings.Fields(lines[i])

    switch v[0] {
    case "acc":
      //increase or decrease accumulator
      arg, err := strconv.Atoi(v[1])
      lib.Check(err)
      accumulator = accumulator + arg
    case "nop":
      //do nothing
    case "jmp":
      //go to new instruction
      arg, err := strconv.Atoi(v[1])
      lib.Check(err)
      i = i + arg -1
    }
		if _, exists := m[i]; exists {
			return accumulator
		} else {
			m[i] = accumulator
		}
    //fmt.Println(v, i, accumulator)
	}
  return accumulator
}
