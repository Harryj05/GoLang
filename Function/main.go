package main
import "fmt"
func simplefunc(){
	fmt.Println("simple function")
}
func add(a, b int) (result int){
	return a+b
}
func multiply(a, b int)(result int){
	return a*b
}
func main(){
	fmt.Println("we are learning function")
	simplefunc()
	ans:= add(3,4)
	fmt.Println("add of two number is: ", ans)
	data:= multiply(3,4)
	fmt.Println("add of two number is: ", data)
}