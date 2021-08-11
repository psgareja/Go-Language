package main
func main(){
	var s string
	s=`hello`
	fmt.Println(s)
	s="<html>\n\t<body>Hello</body?\n</html>"
	fmt.Println(s)
	s=`
	<html>
		<body>
		"hello"
		</body>
	</html>

	
	`
	fmt.Println(s)
}