package main

import (
	"fmt"
	"strconv"
)

func main(){
	
	//path := []string{"R5", "L5", "R5", "R3"}
	//path := []string{"L4", "R2", "R4", "L5", "L3", "L1", "R4", "R5", "R1", "R3", "L3", "L2", "L2", "R5", "R1", "L1", "L2", "R2", "R2", "L5", "R5", "R5", "L2", "R1", "R2", "L2", "L4", "L1", "R5", "R2", "R1", "R1", "L2", "L3", "R2", "L5", "L186", "L5", "L3", "R3", "L5", "R4", "R2", "L5", "R1", "R4", "L1", "L3", "R3", "R1", "L1", "R4", "R2", "L1", "L4", "R5", "L1", "R50", "L4", "R3", "R78", "R4", "R2", "L4", "R3", "L4", "R4", "L1", "R5", "L4", "R1", "L2", "R3", "L2", "R5", "R5", "L4", "L1", "L2", "R185", "L5", "R2", "R1", "L3", "R4", "L5", "R2", "R4", "L3", "R4", "L2", "L5", "R1", "R2", "L2", "L1", "L2", "R2", "L2", "R1", "L5", "L3", "L4", "L3", "L4", "L2", "L5", "L5", "R2", "L3", "L4", "R4", "R4", "R5", "L4", "L2", "R4", "L5", "R3", "R1", "L1", "R3", "L2", "R2", "R1", "R5", "L4", "R5", "L3", "R2", "R3", "R1", "R4", "L4", "R1", "R3", "L5", "L1", "L3", "R2", "R1", "R4", "L4", "R3", "L3", "R3", "R2", "L3", "L3", "R4", "L2", "R4", "L3", "L4", "R5", "R1", "L1", "R5", "R3", "R1", "R3", "R4", "L1", "R4", "R3", "R1", "L5", "L5", "L4", "R4", "R3", "L2", "R1", "R5", "L3", "R4", "R5", "L4", "L5", "R2"}
	path := []string{"R4", "R3", "L3", "L2", "L1", "R1", "L1", "R2", "R3", "L5", "L5", "R4", "L4", "R2", "R4", "L3", "R3", "L3", "R3", "R4", "R2", "L1", "R2", "L3", "L2", "L1", "R3", "R5", "L1", "L4", "R2", "L4", "R3", "R1", "R2", "L5", "R2", "L189", "R5", "L5", "R52", "R3", "L1", "R4", "R5", "R1", "R4", "L1", "L3", "R2", "L2", "L3", "R4", "R3", "L2", "L5", "R4", "R5", "L2", "R2", "L1", "L3", "R3", "L4", "R4", "R5", "L1", "L1", "R3", "L5", "L2", "R76", "R2", "R2", "L1", "L3", "R189", "L3", "L4", "L1", "L3", "R5", "R4", "L1", "R1", "L1", "L1", "R2", "L4", "R2", "L5", "L5", "L5", "R2", "L4", "L5", "R4", "R4", "R5", "L5", "R3", "L1", "L3", "L1", "L1", "L3", "L4", "R5", "L3", "R5", "R3", "R3", "L5", "L5", "R3", "R4", "L3", "R3", "R1", "R3", "R2", "R2", "L1", "R1", "L3", "L3", "L3", "L1", "R2", "L1", "R4", "R4", "L1", "L1", "R3", "R3", "R4", "R1", "L5", "L2", "R2", "R3", "R2", "L3", "R4", "L5", "R1", "R4", "R5", "R4", "L4", "R1", "L3", "R1", "R3", "L2", "L3", "R1", "L2", "R3", "L3", "L1", "L3", "R4", "L4", "L5", "R3", "R5", "R4", "R1", "L2", "R3", "R5", "L5", "L4", "L1", "L1"}
	
	x := 0
	y := 0
	direction := []string{"U","R","D","L"}
	directionInd := 0
	
	for _,move := range path {
		
		rotation := string(move[0])
		distance, err := strconv.Atoi( string(move[1:len(move)]) )
		
		if (err != nil){
			fmt.Println("ERROR")
			return 
		}
		
		if (rotation == "R" && directionInd < 3){
			directionInd++
		}else if (rotation == "R" && directionInd == 3){
			directionInd = 0
		}else if (rotation == "L" && directionInd != 0){
			directionInd--
		}else if (rotation == "L" && directionInd == 0){
			directionInd = 3
		}
		
		if (direction[directionInd] == "U"){
			y = y + distance
		}else if (direction[directionInd] == "R"){
			x = x + distance
		}else if (direction[directionInd] == "D"){
			y = y - distance
		}else if (direction[directionInd] == "L"){
			x = x - distance
		}
		
		//fmt.Println( rotation, distance ,direction[directionInd] , x , y )
		fmt.Print( "(", x,"," , y,")," )	
		//string := move[0]
		// element is the element from someSlice for where we are
	}
	fmt.Println( "Distance equal ", x, y,( y - x ));
}
