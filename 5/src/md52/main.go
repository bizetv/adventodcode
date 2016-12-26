package main

import (
	"fmt"
	"strconv"
	"crypto/md5"
	"bytes"
	"encoding/hex"
	"strings"
)

func main(){
	fmt.Println("Program Start")
	var md5Set = string(`abbhdwsy`)
	//var md5Set = string(`abc`)
	var buffer bytes.Buffer
	var password  = []byte{'_','_','_','_','_','_','_','_'}
	var count = 0;

	//md5Set = "abc"
	
	for i:=0; strings.Contains(string(password), "_");i++{
		buffer.WriteString(md5Set)
		buffer.WriteString(strconv.Itoa(i))
		conc := buffer.String()
		buffer.Reset()
		tmp := []byte(conc)
		tmp2 := md5.Sum( tmp )
		hash := hex.EncodeToString(tmp2[:])

		if (string(hash[0]) == "0" && string(hash[1]) == "0" && string(hash[2]) == "0" && string(hash[3]) == "0" && string(hash[4]) == "0"){
			index, err := strconv.Atoi(string(hash[5]))
			if (err == nil && index >= 0 && index < 8 &&string(password[index]) == "_"){				
				password[index] = hash[6]
				count++
								fmt.Println(hash)
				fmt.Println(string(password))
				fmt.Println(index)

			}
		}
	}
	fmt.Println(password)
}

 
