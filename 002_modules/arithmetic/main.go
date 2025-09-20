package main

import (
    "fmt"
    "module-operation.com/operation"
)

func main(){
    a := 10
    b := 5

    c := operation.Addition(a, b)
    fmt.Printf("Result Addition: %d + %d = %d\n", a, b, c)

    d := operation.Subtraction(a, b)
    fmt.Printf("Result Subtraction: %d - %d = %d\n", a, b, d)

    e := operation.Multiplication(a, b)
    fmt.Printf("Result Multiplicatio: %d * %d = %d\n", a, b, e)

    f, err := operation.Division(a, b)
    if err != nil {
        fmt.Println("Error:", err)
		return
    } else {
        fmt.Printf("Result Devision: %d / %d = %.2f\n", a, b, f)
    }
}

