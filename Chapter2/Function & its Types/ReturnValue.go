package main

import "fmt"
func square(num int) int {
    return num * num
}

func main() {
    result := square(5) 
    fmt.Println("Square of 5:", result)
}
