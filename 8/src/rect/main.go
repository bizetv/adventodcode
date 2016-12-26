package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){
	var input = `rect 1x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 3
rect 1x1
rotate row y=0 by 3
rect 2x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 3
rect 2x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 3
rect 2x1
rotate row y=0 by 5
rect 4x1
rotate row y=0 by 5
rotate column x=0 by 1
rect 4x1
rotate row y=0 by 10
rotate column x=5 by 2
rotate column x=0 by 1
rect 9x1
rotate row y=2 by 5
rotate row y=0 by 5
rotate column x=0 by 1
rect 4x1
rotate row y=2 by 5
rotate row y=0 by 5
rotate column x=0 by 1
rect 4x1
rotate column x=40 by 1
rotate column x=27 by 1
rotate column x=22 by 1
rotate column x=17 by 1
rotate column x=12 by 1
rotate column x=7 by 1
rotate column x=2 by 1
rotate row y=2 by 5
rotate row y=1 by 3
rotate row y=0 by 5
rect 1x3
rotate row y=2 by 10
rotate row y=1 by 7
rotate row y=0 by 2
rotate column x=3 by 2
rotate column x=2 by 1
rotate column x=0 by 1
rect 4x1
rotate row y=2 by 5
rotate row y=1 by 3
rotate row y=0 by 3
rect 1x3
rotate column x=45 by 1
rotate row y=2 by 7
rotate row y=1 by 10
rotate row y=0 by 2
rotate column x=3 by 1
rotate column x=2 by 2
rotate column x=0 by 1
rect 4x1
rotate row y=2 by 13
rotate row y=0 by 5
rotate column x=3 by 1
rotate column x=0 by 1
rect 4x1
rotate row y=3 by 10
rotate row y=2 by 10
rotate row y=0 by 5
rotate column x=3 by 1
rotate column x=2 by 1
rotate column x=0 by 1
rect 4x1
rotate row y=3 by 8
rotate row y=0 by 5
rotate column x=3 by 1
rotate column x=2 by 1
rotate column x=0 by 1
rect 4x1
rotate row y=3 by 17
rotate row y=2 by 20
rotate row y=0 by 15
rotate column x=13 by 1
rotate column x=12 by 3
rotate column x=10 by 1
rotate column x=8 by 1
rotate column x=7 by 2
rotate column x=6 by 1
rotate column x=5 by 1
rotate column x=3 by 1
rotate column x=2 by 2
rotate column x=0 by 1
rect 14x1
rotate row y=1 by 47
rotate column x=9 by 1
rotate column x=4 by 1
rotate row y=3 by 3
rotate row y=2 by 10
rotate row y=1 by 8
rotate row y=0 by 5
rotate column x=2 by 2
rotate column x=0 by 2
rect 3x2
rotate row y=3 by 12
rotate row y=2 by 10
rotate row y=0 by 10
rotate column x=8 by 1
rotate column x=7 by 3
rotate column x=5 by 1
rotate column x=3 by 1
rotate column x=2 by 1
rotate column x=1 by 1
rotate column x=0 by 1
rect 9x1
rotate row y=0 by 20
rotate column x=46 by 1
rotate row y=4 by 17
rotate row y=3 by 10
rotate row y=2 by 10
rotate row y=1 by 5
rotate column x=8 by 1
rotate column x=7 by 1
rotate column x=6 by 1
rotate column x=5 by 1
rotate column x=3 by 1
rotate column x=2 by 2
rotate column x=1 by 1
rotate column x=0 by 1
rect 9x1
rotate column x=32 by 4
rotate row y=4 by 33
rotate row y=3 by 5
rotate row y=2 by 15
rotate row y=0 by 15
rotate column x=13 by 1
rotate column x=12 by 3
rotate column x=10 by 1
rotate column x=8 by 1
rotate column x=7 by 2
rotate column x=6 by 1
rotate column x=5 by 1
rotate column x=3 by 1
rotate column x=2 by 1
rotate column x=1 by 1
rotate column x=0 by 1
rect 14x1
rotate column x=39 by 3
rotate column x=35 by 4
rotate column x=20 by 4
rotate column x=19 by 3
rotate column x=10 by 4
rotate column x=9 by 3
rotate column x=8 by 3
rotate column x=5 by 4
rotate column x=4 by 3
rotate row y=5 by 5
rotate row y=4 by 5
rotate row y=3 by 33
rotate row y=1 by 30
rotate column x=48 by 1
rotate column x=47 by 5
rotate column x=46 by 5
rotate column x=45 by 1
rotate column x=43 by 1
rotate column x=38 by 3
rotate column x=37 by 3
rotate column x=36 by 5
rotate column x=35 by 1
rotate column x=33 by 1
rotate column x=32 by 5
rotate column x=31 by 5
rotate column x=30 by 1
rotate column x=23 by 4
rotate column x=22 by 3
rotate column x=21 by 3
rotate column x=20 by 1
rotate column x=12 by 2
rotate column x=11 by 2
rotate column x=3 by 5
rotate column x=2 by 5
rotate column x=1 by 3
rotate column x=0 by 4`	

	//input = `abba[mnop]qrst`
	fmt.Println("Program Start")
	
	var inputArray = strings.Split(input, "\n")
	var mapArray  = [6][50] string{ { ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", "#"},
		{ ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", "."},
		{ ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", "."},
		{ ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", "."},
		{ ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", "."},
		{ ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", ".", ".", ".", ".",".", ".",".", ".", ".",".", "."}}
	for _, row := range inputArray{
		instructArr := strings.Split(row, ` `)
		fmt.Println(instructArr)
		
		switch instructArr[0]{
		case "rect":
			Rect(instructArr[1],&mapArray)
			break
		case "rotate":
			Rotate(instructArr[1], instructArr[2], instructArr[4], &mapArray)
			break
			
		}
		fmt.Println(mapArray[0])
		fmt.Println(mapArray[1])
		fmt.Println(mapArray[2])
		fmt.Println(mapArray[3])
		fmt.Println(mapArray[4])
		fmt.Println(mapArray[5])
	}
	
	var count = 0
	for y2 := 0; y2 < 6; y2++{
		for x := 0; x < 50; x++{
			if (mapArray[y2][x] == "#"){
				count++
			}			
		}
	}
	fmt.Println(count)
}

func Rect(size string, mapArray *[6][50]string){
	sizeArr := strings.Split(size, "x")

	xMax,_ := strconv.Atoi(sizeArr[0])
	yMax,_ := strconv.Atoi(sizeArr[1])
	for y := 0; y < yMax; y++{	
		for x := 0; x < xMax; x++{
			mapArray[y][x] = "#"
		}
	}
}
func Rotate(direction, from , to string, mapArray *[6][50] string){
	y := 0
	x := 0
	toInt, _ := strconv.Atoi(to)
	if (direction == "row"){
		fromTmp := strings.Split(from, "=")
		y, _= strconv.Atoi(fromTmp[1])
		var tmpRow =  mapArray[y]
		for i := 0; i < 50; i++{
			if (i < toInt){
				mapArray[y][i] = tmpRow[len(tmpRow)-toInt+i]
			}else{
				mapArray[y][i] = tmpRow[i-toInt]
			}
		}
	}else if (direction == "column"){
		fromTmp := strings.Split(from, "=")
		x, _= strconv.Atoi(fromTmp[1])
		var tmpColumn = [6] string{mapArray[0][x],mapArray[1][x],mapArray[2][x],mapArray[3][x],mapArray[4][x],mapArray[5][x]}
		for i := 0; i < 6; i++{
			if (i < toInt){
				mapArray[i][x] = tmpColumn[len(tmpColumn) -toInt+i]
			}else{
				mapArray[i][x] = tmpColumn[i-toInt]
			}
		}
	}

	
}
