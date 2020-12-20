package main

import (
	"bufio"
    "fmt"
    "os"
	"strings"
	"regexp"
)

func readFile(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func checkPassports(lines []string){
	numOfValidPassports := 0
	var requiredFields = [7]string{"byr:", "iyr:", "eyr:", "hcl:", "hgt:", "ecl:", "pid:"}
	concatenatedPassport := ""
	for _, line := range lines{
	checked := 0

		if line != ""{
			concatenatedPassport = concatenatedPassport + " " + line
			checked = 0
		}else{
			for _, field := range requiredFields{
				if strings.Contains(concatenatedPassport, field){
					checked += 1
				}else{
					concatenatedPassport = ""
					break
				}
			}
			if checked == 7{
				checked = 0
				concatenatedPassport = concatenatedPassport + " "
				var validbyr = regexp.MustCompile(`byr:((19[2-9]\d( |\n))|(200[0-2]( |\n)))`)
				var validiyr = regexp.MustCompile(`iyr:((2020( |\n))|(201[0-9]( |\n)))`)
				var valideyr = regexp.MustCompile(`eyr:((2030( |\n))|(202[0-9]( |\n)))`)
				var validhcl = regexp.MustCompile(`hcl:#([0-9a-f]{6})( |\n)`)
				var validhgt = regexp.MustCompile(`hgt:(((1[5-8][0-9])cm( |\n)|(19[0-3])cm( |\n))|((59in( |\n))|(6[0-9]in( |\n))|(7[0-6]in( |\n))))`)
				var validecl = regexp.MustCompile(`ecl:(amb( |\n)|blu( |\n)|brn( |\n)|gry( |\n)|grn( |\n)|hzl( |\n)|oth( |\n))`)
				var validpid = regexp.MustCompile(`pid:([0-9]{9})( |\n)`)
				
				if validbyr.MatchString(concatenatedPassport) && validiyr.MatchString(concatenatedPassport) &&
				valideyr.MatchString(concatenatedPassport) && validhcl.MatchString(concatenatedPassport) &&
				validhgt.MatchString(concatenatedPassport) && validecl.MatchString(concatenatedPassport) &&
				validpid.MatchString(concatenatedPassport){
					numOfValidPassports += 1
					fmt.Println(numOfValidPassports)
				}
				
				concatenatedPassport = ""

			}
		}	
	}
}

func main(){
	lines, err := readFile("input.txt")
	checkPassports(lines)
	fmt.Println(err)
}