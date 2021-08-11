import("fmt"; "unicode/utf8")
func main(){
	name:='Italic'
	fmt.Println(len(name))
	Println(utf8.RuneCountInString(name))


}