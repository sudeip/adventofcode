package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "sort"
)

type BoardingPass struct {
  seatId int
  row int
  col int
	seat string
}

type manifest []BoardingPass

func (m manifest) Len() int {
    return len(m)
}

func (m manifest) Less(i, j int) bool {
    return m[i].seatId < m[j].seatId
}

func (m manifest) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func main() {
  //Read the file containing input
  file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  var seats []string
  for scanner.Scan() {
      seats = append(seats, scanner.Text())
  }
  file.Close()
  var boardingPassList = []BoardingPass{}
  for i:=0; i<len(seats); i++{
    var boardingPass = new(BoardingPass)
    rowId, colId, rowMin, colMin, rowMax, colMax := 0, 0, 0, 0, 127, 7
    var seatStr = seats[i]
    for j:=0; j<len(seatStr); j++ {
      switch j {
        case 0:
          if(seatStr[j] == 'F'){
            rowMax = rowMax - 64
          }else{
            rowMin = rowMin + 64
          }
        case 1:
          if(seatStr[j] == 'F'){
            rowMax = rowMax - 32
          } else {
            rowMin = rowMin + 32
          }
        case 2:
          if(seatStr[j] == 'F'){
            rowMax = rowMax - 16
          } else {
            rowMin = rowMin + 16
          }
        case 3:
          if(seatStr[j] == 'F'){
            rowMax = rowMax - 8
          } else {
            rowMin = rowMin + 8
          }
        case 4:
          if(seatStr[j] == 'F'){
            rowMax = rowMax - 4
          } else {
            rowMin = rowMin + 4
          }
        case 5:
          if(seatStr[j] == 'F'){
            rowMax = rowMax - 2
          } else {
            rowMin = rowMin + 2
          }
        case 6:
          if(seatStr[j] == 'F'){
            rowId = rowMax - 1
          } else {
            rowId = rowMin + 1
          }
        case 7:
          if(seatStr[j] == 'L'){
            colMax = colMax - 4
          } else {
            colMin = colMin + 4
          }
        case 8:
          if(seatStr[j] == 'L'){
            colMax = colMax - 2
          } else {
            colMin = colMin + 2
          }
        case 9:
          if(seatStr[j] == 'L'){
            colId = colMax - 1
          } else {
            colId = colMin + 1
          }
        }
      }
      //setup struct
      boardingPass.seat = seatStr
      boardingPass.row = rowId
      boardingPass.col = colId
      boardingPass.seatId = rowId*8+colId
      boardingPassList = append(boardingPassList, *boardingPass)
    }
  sort.Sort(manifest(boardingPassList))
  fmt.Println("part1 ans: ", boardingPassList[len(boardingPassList)-1].seatId)

  for i := 1; i < len(boardingPassList); i++ {
    if(boardingPassList[i].seatId-boardingPassList[i-1].seatId == 2){
      fmt.Println("part2 ans: ", boardingPassList[i].seatId-1)
      break
    }
  }
}
