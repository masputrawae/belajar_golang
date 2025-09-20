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

func Division(a, b int) (float64, error) {
    if b == 0 {
        return 0, errors.New("Cannot be divided by 0")
    }
    return float64(a) / float64(b), nil
}

