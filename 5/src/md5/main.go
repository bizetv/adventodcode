package main

import (
	"fmt"
	"strconv"
	"crypto/md5"
	"bytes"
	"encoding/hex"
)

func main(){
	fmt.Println("Program Start")
	var md5Set = string(`abbhdwsy`)
	var buffer bytes.Buffer
	var password = ""

	//md5Set = "abc"
	
	for i:=0; len(password) < 8;i++{
		buffer.WriteString(md5Set)
		buffer.WriteString(strconv.Itoa(i))
		conc := buffer.String()
		buffer.Reset()
		tmp := []byte(conc)
		tmp2 := md5.Sum( tmp )
		hash := hex.EncodeToString(tmp2[:])

		if (string(hash[0]) == "0" && string(hash[1]) == "0" && string(hash[2]) == "0" && string(hash[3]) == "0" && string(hash[4]) == "0"){
			password += string(hash[5])
			fmt.Println(hash)
		}
	}
	fmt.Println(password)
}

 
