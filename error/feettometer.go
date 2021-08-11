func main(){
    arg:=os.Args[1]
    feet,_err:=strconv.ParseFloat(arg,64)
    
    if err!=nil{
        fmt.Println("Error",_err)
        return 
    }
    else{
        meter:=feet*0.3048
    fmt.Println("%g feet is % meter",feet,meter)

    }
}