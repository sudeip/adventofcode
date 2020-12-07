package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
  file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  var lines []string
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

	groups := strings.Split(strings.Join(lines, "\n"), "\n\n")

	var count = 0
	for _, group := range groups {
		count += matchEvery(group)
	}
	fmt.Println(count)
	//fmt.Println(len(customs))
}

func matchEvery(input string) int {
	lines := strings.Split(input, "\n")
	m := map[rune]int{}
	for _, line := range lines {
		for _, question := range line {
			if _, exists := m[question]; exists {
				m[question]++
			} else {
				m[question] = 1
			}
		}
	}
  //fmt.Println(m)
	count := 0
	for _, v := range m {
		if v == len(lines) {
			count++
		}
	}
	return count
}
