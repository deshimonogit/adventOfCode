package main

import (
    "fmt"
    "io"
    "os"
	"sort"
)

func readFile(filePath string) (numbers []int){
	//Loads all lines to array as integers
    file, err := os.Open(filePath)
    if err != nil{
        fmt.Println(err)
    }
    var line int
    for{
        _, err := fmt.Fscanf(file, "%d\n", &line)
        if err != nil{
            fmt.Println(err)
            if err == io.EOF{
                return
            }
            panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))

        }
        numbers = append(numbers, line)
    }
	defer file.Close()
    return numbers
}

func find2020andMultiply(nrs []int){
	sort.Ints(nrs)
	left := 0
	right := len(nrs) - 1
	nextToLeft := 1
	for left < right{
		if nrs[left] + nrs[nextToLeft] + nrs[right] == 2020{
			fmt.Println(nrs[left]*nrs[nextToLeft]*nrs[right])
			break
		}else if nrs[left] + nrs[nextToLeft] + nrs[right] < 2020{
			left += 1
		}else{
			right -= 1
		}
	}
}

func main(){
    numbers := readFile("input.txt")
	find2020andMultiply(numbers)
}