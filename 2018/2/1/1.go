package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func occurCount(id string) (bool, bool) {
	two, three := false, false
	var alphaArray = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	for _, alpha := range alphaArray {
		alphaCount := strings.Count(id, alpha)
		if alphaCount == 2 {
			two = true
		} else if alphaCount == 3 {
			three = true
		}
		if two == true && three == true {
			break
		}
	}
	return two, three
}

func main() {
	file, err := os.Open("input.txt")
	two, three := 0, 0

	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			file.Seek(0, 0)
			break
		}
		isTwo, isThree := occurCount(string(line))
		if isTwo {
			two++
		}
		if isThree {
			three++
		}
	}
	fmt.Println(two * three)
}
