package operation

import (
    "errors"
)

func Addition(a, b int) int {
    return a + b
}

func Subtraction(a, b int) int {
    return a - b
}

func Multiplication(a, b int) int {
    return a * b
}

func Division(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("Cannot be divided by 0")
    }
    return a / b, nil
}

