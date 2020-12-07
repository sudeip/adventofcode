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

	var customs []string

	for i:=0;i<len(lines);i++ {
		if(i >0 && lines[i] != ""){
			customs[len(customs)-1] = unique(customs[len(customs)-1]+lines[i])
			//fmt.Println(len(customs[len(customs)-1]))
		} else {
			customs = append(customs, unique(lines[i]))
		}
  }
	var length = 0
	for _, customForm := range customs {
		length = length + len(customForm)
	}

	fmt.Println(length)
	fmt.Println(len(customs))
}

func unique(input string) string {
    occured := map[byte]bool{}
    result := []string{}
    for e:=0;e<len(input);e++ {
        // check if already the mapped
        // variable is set to true or not
        if occured[input[e]] != true {
            occured[input[e]] = true
            // Append to result slice.
            result = append(result, string(input[e]))
        }
    }
    return strings.Join(result,"")
}
