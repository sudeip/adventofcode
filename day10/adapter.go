package main

import (
	"fmt"
	"../lib"
  "strconv"
  "sort"
)
var cache = map[int]int64{}

func main() {
	lines := lib.ReadInput()

	ints := []int{0}
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		lib.Check(err)
		ints = append(ints, i)
	}
  fmt.Println(part1(ints))
  fmt.Println(part2(ints))
}

func part1(jolts []int) int{
  sort.Ints(jolts)
  //and 1 for your device that has +3 jolts
  count1, count3 := 0, 1
  //fmt.Println(jolts)
  for i := 1; i < len(jolts); i++ {
    if(jolts[i]-jolts[i-1] == 1){
      count1++
    } else if(jolts[i]-jolts[i-1] == 3){
      count3++
    }else{
  		panic("Oh no! the chain is broke ... Bad input")
    }
  }
  return count1*count3
}

func part2(jolts []int) int64 {
  arrSize := len(jolts)
	if arrSize <= 2 {
		return 1
	}
	l := arrSize - 1
	last := jolts[l]
	res, ok := cache[last]
	if ok {
		return res
	}
	if last-jolts[l-1] <= 3 {
		res += part2(jolts[:l])
	}
	if arrSize >= 3 && last-jolts[l-2] <= 3 {
		res += part2(jolts[:l-1])
	}
	if arrSize >= 4 && last-jolts[l-3] <= 3 {
		res += part2(jolts[:l-2])
	}
	cache[last] = res
	return res
}

// func part2(jolts []int) int{
//   sort.Ints(jolts)
//   //count := 0
//   //m := make(map[string]bool)
//   var arr = []string{}
//
//   deviceJolt := jolts[len(jolts)-1]+3
//   //add fisrt for test
//   //m[arrayToString(jolts)] = true
//   fmt.Println(arr)
//   var ignore = []int
//   for i := 1; i < len(jolts); i++ {
//     if(jolts[i]-jolts[i-1]==1){
//       arr = append(arr, arrayToString(jolts, ignore))
//     }
//   }
//
//   return deviceJolt
// }
//
// func arrayToString(a []int) string {
//     b := make([]string, len(a))
//     for i, v := range a {
//         b[i] = strconv.Itoa(v)
//     }
//     return strings.Join(b,"-")
// }

// func arrayToString(A []int) string {
//     var buffer bytes.Buffer
//     for i := 0; i < len(A); i++ {
//         buffer.WriteString(strconv.Itoa(A[i]))
//     }
//     return buffer.String()
// }


//func arrayToString(a []int) string {
    // b := ""
    // for _, v := range a {
    //     if len(b) > 0 {
    //         b += ","
    //     }
    //     b += strconv.Itoa(v)
    // }
    //
    // return b
//}
