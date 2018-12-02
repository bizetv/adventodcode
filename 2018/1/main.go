package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func contains(s []int, e int) bool {
	//fmt.Println(len(s))
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	frequency := 0
	//i := 0
	var frequencyArray []int
	frequencyArray = append(frequencyArray, frequency)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {

			file.Seek(0, 0)
			continue
		}
		modifValue, err := strconv.Atoi(string(line))
		check(err)
		//fmt.Println(" frquency[" + strconv.Itoa(frequency) + "] modifValue[" + strconv.Itoa(modifValue) + "]")
		frequency += modifValue
		if contains(frequencyArray, frequency) {
			fmt.Println(frequency)
			break
		}
		frequencyArray = append(frequencyArray, frequency)
		//fmt.Println(frequencyArray)

	}
	fmt.Println(frequency)
}
