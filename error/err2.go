func main(){
	age:=os.Args[1]
	n,err=strconv.Atoi(age)
	if err!=nil{
		fmt.Println("No Error",Error)
		return 
	}
	else{
		fmt.Println(n)
	}


}