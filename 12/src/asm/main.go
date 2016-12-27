package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){
	var input = `cpy 1 a
cpy 1 b
cpy 26 d
jnz c 2
jnz 1 5
cpy 7 c
inc d
dec c
jnz c -2
cpy a c
inc a
dec b
jnz b -2
cpy c b
dec d
jnz d -6
cpy 14 c
cpy 14 d
inc a
dec d
jnz d -2
dec c
jnz c -5`

	/*input =`cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`*/
	var inputArray = strings.Split(input, "\n")
	//input = "A(1x5)BC"
	//input =  "X(8x2)(3x3)ABCY"
	mapVar := make(map[string]int)
	mapVar["a"] = 0
	mapVar["b"] = 0
	mapVar["c"] = 0
	mapVar["d"] = 0
	
	for i := 0; i < len(inputArray); {
		instructArr := strings.Split(inputArray[i], ` `)
		switch instructArr[0]{
		case "cpy":
			var tmpVal = 0
			if (instructArr[1] == "a" || instructArr[1] == "b" || instructArr[1] == "c" || instructArr[1] == "d"){
				tmpVal = mapVar[instructArr[1]]
			}else{
				tmpVal, _ = strconv.Atoi(instructArr[1])
			}
			mapVar[instructArr[2]] = tmpVal
			i++
			break
		case "inc":
			mapVar[instructArr[1]]++
			i++
			break
		case "dec":
			mapVar[instructArr[1]]--
			i++
			break
		case "jnz":
			var tmpVal = 0
			if (instructArr[1] == "a" || instructArr[1] == "b" || instructArr[1] == "c" || instructArr[1] == "d"){
				tmpVal = mapVar[instructArr[1]]
			}else{
				tmpVal, _ = strconv.Atoi(instructArr[1])
			}
			if (tmpVal != 0){
				if (instructArr[2] == "a" || instructArr[2] == "b" || instructArr[2] == "c" || instructArr[2] == "d"){
					tmpVal = mapVar[instructArr[2]]
				}else{	
					tmpVal, _ = strconv.Atoi(instructArr[2])
				}
				i = i + tmpVal
			}else{
				i++
			}
			break
		}
		//fmt.Println(i , instructArr)
		//fmt.Println(mapVar)
	}
	fmt.Println(mapVar["a"])
	
}
