package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)


func main() {
    scanner := bufio.NewScanner(os.Stdin)
    table := make([][]string, 0)
    for scanner.Scan() {
        row := strings.Split(scanner.Text(), "\t")
        table = append(table, row)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)

    } else {
        sum := 0
        val := 0
        for _, row := range table {
            for i, valA := range row[:len(row) - 1] {
                a, _ := strconv.Atoi(valA)
                for _, valB := range row[i + 1:] {
                    b, _ := strconv.Atoi(valB)
                    if a % b == 0 {
                        val = a / b
                    } else if b % a == 0 {
                        val = b / a
                    }
                    if val > 0 {
                        break
                    }
                }
                if val > 0 {
                    break
                }
            }
            sum += val
            val = 0
        }
        fmt.Println(sum)
    }
}
