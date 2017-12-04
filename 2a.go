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
        for _, row := range table {
            min_val, _ := strconv.Atoi(row[0])
            max_val, _ := strconv.Atoi(row[0])
            for _, valIn := range row {
                val, _ := strconv.Atoi(valIn)
                if val > max_val {
                    max_val = val
                } else if val < min_val {
                    min_val = val
                }
            }
            sum += max_val - min_val
        }
        fmt.Println(sum)
    }
}
