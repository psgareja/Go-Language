package main
func main(){
	msg:=os.Args[1]
	l:=len(msg)
	s:=msg+strings.Repeat("!",l)

}