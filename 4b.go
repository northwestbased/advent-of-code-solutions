package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "sort"
)

func sortWord(a string) string {
    ints :=  []int{}
    for _, i := range a {
      ints = append(ints, int(i))
    }
    sort.Ints(ints)

    letters := ""
    for _, i := range ints {
        letters += string(i)
    }

    return letters
}

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
                mainWord := sortWord(word)
                if m[mainWord] == 1 {
                    flag = 1
                    break
                }
                m[mainWord] = 1
            }
            if flag == 0 {
                total += 1
            }
        }
        fmt.Println(total)
    }
}
