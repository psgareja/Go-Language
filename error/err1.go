package main

import (
	"fmt"
	"strconv"
	"os"
)
func main(){
	n,err:=strconv.Atoi(os.Args[1])
	fmt.Println(n)
	fmt.Println(err)
	
}