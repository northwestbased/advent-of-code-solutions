package main


import (
    "fmt"
    "math"
)

const input float64 = 368078

func main() {
    var side_length float64 = math.Ceil(math.Sqrt(input))
    var center float64 = math.Ceil(side_length / 2.0)
    var difference float64 = float64(int(input) % int(side_length))
    var offset = math.Abs(center - difference - 1)
    fmt.Println(offset + center)
}
