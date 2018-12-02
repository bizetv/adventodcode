package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func SliceCompare(word1, word2 string) string {
	if len(word1) != len(word2) {
		return ""
	}
	for i := 0; i < len(word1); i++ {
		a := word1[:i] + word1[i+1:]
		b := word2[:i] + word2[i+1:]

		if a == b {
			return a
		}
	}
	return ""
}

func part2(input []string) string {
	for _, word1 := range input {
		for _, word2 := range input {
			if word1 == word2 {
				continue
			}
			if res := SliceCompare(word1, word2); res != "" {
				return res
			}
		}
	}
	return ""
}

func main() {
	file, err := os.Open("input.txt")

	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	b, err := ioutil.ReadAll(reader)
	check(err)
	input := string(b)
	inputArray := strings.Split(input, "\n")
	//fmt.Println(len(inputArray))
	fmt.Println(part2(inputArray))

}
