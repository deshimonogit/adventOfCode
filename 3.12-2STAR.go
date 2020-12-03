package main

import (
    "fmt"
    "io"
    "os"
)

func readFile(filePath string) (strings []string){
	//Loads all lines to array as strings
    file, err := os.Open(filePath)
    if err != nil{
        fmt.Println(err)
    }
    var line string
    for{
        _, err := fmt.Fscanf(file, "%s\n", &line)
        if err != nil{
            fmt.Println(err)
            if err == io.EOF{
                return
            }
            panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))

        }
        strings = append(strings, line)
    }
	defer file.Close()
    return strings
}

func letsSlide(worldMap []string, test []int){
	posX := 0
	treeCounter := 0
	movX := test[0]
	movY := test[1]
	if movY == 1{
		for _, row := range worldMap{
			if string(row[posX]) == "#"{
				treeCounter += 1
			}
			if posX < 30 - (movX - 1){
				posX += movX
			}else{
				alreadyMoved := 30 - posX
				posX = movX - 1 - alreadyMoved
			}
		}
	}else{
		for i, row := range worldMap{
			if i == 0 || i%2 == 0{
				if string(row[posX]) == "#"{
					treeCounter += 1
				}
				if posX < 30 - (movX - 1){
					posX += movX
				}else{
					alreadyMoved := 30 - posX
					posX = movX - 1 - alreadyMoved
				}
			}
		}
	}
	fmt.Println(treeCounter)
}

func main(){
    lines := readFile("input.txt")
	testOne := []int{1,1}
	testTwo := []int{3,1}
	testThree := []int{5,1}
	testFour := []int{7,1}
	testFive := []int{1,2}
	letsSlide(lines, testOne)
	letsSlide(lines, testTwo)
	letsSlide(lines, testThree)
	letsSlide(lines, testFour)
	letsSlide(lines, testFive)
}