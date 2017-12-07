package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)


func main() {
    scanner := bufio.NewScanner(os.Stdin)
    passphrases := make([][]string, 0)
    for scanner.Scan() {
        words := strings.Split(scanner.Text(), " ")
        passphrases = append(passphrases, words)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    } else {
        total := 0
        for _, passphrase := range passphrases {
            m := make(map[string]int)
            flag := 0
            for _, word := range passphrase {
                if m[word] == 1 {
                    flag = 1
                    break
                }
                m[word] = 1
            }
            if flag == 0 {
                total += 1
            }
        }
        fmt.Println(total)
    }
}
