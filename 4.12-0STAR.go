package main

import (
	"bufio"
    "fmt"
    "os"
	"strings"
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
					
					//fmt.Println("Found ", field)
					checked += 1
				}else{
					fmt.Println("NOT FOUND")
					break
				}
			}
			if checked == 7{
				checked = 0
				fmt.Println("Concatenated: ", concatenatedPassport)
				concatenatedPassport = ""
				numOfValidPassports += 1
				fmt.Println(numOfValidPassports)
			}
		}	
	}
}

func main(){
	lines, err := readFile("input.txt")
	checkPassports(lines)
	fmt.Println(err)
}