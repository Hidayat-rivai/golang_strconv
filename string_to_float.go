package main

import "fmt"
import "strconv"

func main(){
	var str = "30.7"
	var flt, err = strconv.ParseFloat(str,8)
	
	if err == nil {
		fmt.Println(flt)
	}
}