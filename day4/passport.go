package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"image/color"
)

type Passport struct {
   byr int
   iyr int
   eyr int
	 hgt string
   hcl string
   ecl string
   pid string
   cid string
}

var passports = []Passport{}

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

	var documents []string

	for i:=0;i<len(lines);i++ {
		if(i >0 && lines[i] != ""){
			documents[len(documents)-1] = documents[len(documents)-1]+ " "  + lines[i]
			passport := passports[len(passports)-1]
			passports[len(passports)-1] = createUpdatePassport(passport, lines[i])
			//fmt.Println(passports)

		} else {
			documents = append(documents, lines[i])
			var passport = new(Passport)

			passports=append(passports, createUpdatePassport(*passport, lines[i]))
			//fmt.Println(passports)

		}
  }
	var present = 0
	var valid = 0
	for _, document := range documents {
		if(isPresent(document)){
			present++
		}
		//fmt.Println(document)
	}
	for _, passport := range passports {
		//fmt.Println(passport.hgt)

		if(isValid(passport)){
			valid++
		}
	}
	//fmt.Println(passports)
	fmt.Println(len(documents))
	fmt.Println(present)
	fmt.Println(valid)
}

func createUpdatePassport(passport Passport, input string) (Passport){
	if input != "" {
		s := strings.Fields(input)
		for _, v := range s {
			field := strings.Split(v,":")
			switch field[0] {
			case "byr":
				i1, err := strconv.Atoi(field[1])
				if err == nil {
					passport.byr = i1
				}
			case "iyr":
				i1, err := strconv.Atoi(field[1])
				if err == nil {
					passport.iyr = i1
				}
			case "eyr":
				i1, err := strconv.Atoi(field[1])
				if err == nil {
					passport.eyr = i1
				}
			case "hgt":
				passport.hgt = field[1]
			case "hcl":
				passport.hcl = field[1]
			case "ecl":
				passport.ecl = field[1]
			case "pid":
				passport.pid = field[1]
			case "cid":
				passport.cid = field[1]
			}
		}
	}
	return passport
}

func isValid(passport Passport)(bool){
	if(passport.byr == 0 || passport.iyr == 0 || passport.eyr == 0 || passport.hgt == "" || passport.hcl == "" || passport.ecl == "" || passport.pid == "" ){
		return false
	}
	if(passport.byr < 1920 || passport.byr > 2002){
		return false
	}
	if(passport.iyr < 2010 || passport.iyr > 2020){
		return false
	}
	if(passport.eyr < 2020 || passport.eyr > 2030){
		return false
	}
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	if(strings.HasSuffix(passport.hgt, "cm")){
		i1, err := strconv.Atoi(strings.TrimRight(passport.hgt, "cm"))
		if (err == nil && (i1 <150 || i1 > 193 )) {
			return false
		}
	} else if(strings.HasSuffix(passport.hgt, "in")){
		i1, err := strconv.Atoi(strings.TrimRight(passport.hgt, "in"))
		if (err == nil && (i1 <59 || i1 > 76 )) {
			return false
		}
	} else {
		return false
	}

	//(Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	c, err := ParseHexColor(passport.hcl)
	if (err != nil){
		fmt.Println(c)
		return false
	}
	//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if(!strings.Contains("amb blu brn gry grn hzl oth", passport.ecl)){
		return false
	}
	//pid (Passport ID) - a nine-digit number, including leading zeroes.
	if(len(passport.pid) == 9){
		i1, err := strconv.Atoi(passport.pid)
		if (err != nil){
			fmt.Print(i1)
			return false
		}
	} else {
		return false
	}

	return true
}

func ParseHexColor(s string) (c color.RGBA, err error) {
    c.A = 0xff
    switch len(s) {
    case 7:
        _, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
    //case 4:
      //  _, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
        // Double the hex digits:
      //  c.R *= 17
      //  c.G *= 17
      //  c.B *= 17
    default:
        err = fmt.Errorf("invalid length, must be 7 or 4")
    }
    return
}

func isPresent(input string)(bool){
	reqFields := []string{"byr:","iyr:","eyr:","hgt:","hcl:","ecl:","pid:"}
	//optField := "cid:"
	var found = true
	for _, reqField := range reqFields {
		if(!strings.Contains(input, reqField)){
			found = false
		}
	}
	return found
}
