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
        for i := 0; i < len(input) - 1; i++ {
            a := input[i]
            b := input[i + 1]
            if a == b {
                output += int(a - '0')
            }
        }
        if input[0] == input[len(input) - 1] {
          output += int(input[0] - '0')
        }
    }

    fmt.Println(output)
}
