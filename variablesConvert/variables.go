package main

import (
	"fmt"
	"reflect"
)



func main() {
    var x int = 42
    f:=float64(x)
    result := f/9;
    sameResult:= float64(x) / 9
    // f := float64(x/9)
    
    fmt.Printf("x is of type: %s\n", reflect.TypeOf(x))
    // Output: x is of type: int
    fmt.Println(x)
    
    fmt.Printf("f is of type: %s\n", reflect.TypeOf(f))
    fmt.Println(f)

    fmt.Println(result)
    fmt.Println(sameResult)
    // Output: f is of type: float64
}