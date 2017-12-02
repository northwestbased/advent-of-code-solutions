package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var input string = ""
    var output int = 0
    for scanner.Scan() {
        input += scanner.Text()
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)

    } else {
        length := len(input)
        println("length", length)
        for i := 0; i < length; i++ {
            a := input[i]
            b_index :=  (i + length / 2) % length
            b := input[b_index]
            if a == b {
                output += int(a - '0')
            }
        }
    }

    fmt.Println(output)
}
