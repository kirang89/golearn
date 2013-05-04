package main

import (
    "fmt"
    "math"
)

type errorstring struct {
    message string
}

func (e *errorstring) Error() string{
    return e.message
}

func NegativeSqrtError() error {
    return &errorstring{"Cannot find square root of a negative number"}
}

func Sqrt(f float64) (float64, error) {
    if f < 0{
        return 0, NegativeSqrtError()
    }
    return math.Sqrt(f), nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
