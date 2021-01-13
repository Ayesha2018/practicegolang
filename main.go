package main
import "fmt"
func main() {
    fmt.Println("Go testing tut0rial")
    result := calculate(2)
    fmt.Println(result)
}
func calculate(x int)(result int){    
    result = x + 2
    return result  
}