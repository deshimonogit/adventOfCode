package main

import (
    "fmt"
    "io"
    "os"
	"strings"
	"strconv"
)

func readFile(filePath string) (strings []string){
	//Loads all lines to array as strings
    file, err := os.Open(filePath)
    if err != nil{
        fmt.Println(err)
    }
    var line string
	var sign string
	var password string
    for{
        _, err := fmt.Fscanf(file, "%s %s %s\n", &line, &sign, &password)
        if err != nil{
            fmt.Println(err)
            if err == io.EOF{
                return
            }
            panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))

        }
        strings = append(strings, line, sign, password)
    }
	defer file.Close()
    return strings
}

func checkPasswords(lines []string){
	numOfValidPasswords := 0
	
	for i := 0; i<len(lines); i++{
		if i == 0 || i%3 == 0{
			controlSum := 0
			pwRange := strings.Split(lines[i], "-")
			pwMin, _ := strconv.ParseInt(pwRange[0], 10, 64)
			pwMax, _ := strconv.ParseInt(pwRange[1], 10, 64)
			pwLetter := string(lines[i+1][0])
			pwPassword := lines[i+2]
			
			if string(pwPassword[pwMin-1]) == pwLetter{
				controlSum+=1
			}
			
			if string(pwPassword[pwMax-1]) == pwLetter{
				controlSum+=1
			}
			
			if controlSum == 1{
			numOfValidPasswords+=1
			}
		}
	}
	
	fmt.Println(numOfValidPasswords)
}

func main(){
    lines := readFile("input.txt")
	checkPasswords(lines)
	
}