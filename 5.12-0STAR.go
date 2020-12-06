package main

import (
	"bufio"
    "fmt"
    "os"
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

func checkBoardPass(lines []string){
	/*
	minRows = 0
	maxRows = 127
	minColumns
	maxColumns = 8
	highestSeatID := 0
	*/
	
	for _, line := range lines{
		rowID := 0
		colID := 0
		row := line[0:6]
		column := line[6:9]
		
		for _, rowElement := range row{
			if rowElement == "F"{
				maxRows = maxRows /2
			}else{
				minRows = maxRows /2
			}
		}
		
		for _, rowElement := range row{
		
		}
		fmt.Println(row, column)
	}
}

func main(){
	lines, err := readFile("input_two.txt")
	checkBoardPass(lines)
	fmt.Println(err)
}