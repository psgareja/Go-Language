package main
import "fmt"
func main(){
	args:=or.Args
	if len(args)!=3{
		fmt.Println("Usage:[username] [Password]")
		return 

	}
	