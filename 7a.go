package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)


func main() {
    scanner := bufio.NewScanner(os.Stdin)

    possibleRoot := []string{}
    children := make(map[string]bool)

    for scanner.Scan() {
        line := scanner.Text()
        splitLine := strings.Split(line, " ")


        if len(splitLine) > 2 {
            possibleRoot = append(possibleRoot, splitLine[0])
            splitLine =  splitLine[3:len(splitLine)]

            for _, word := range splitLine {
                if word[len(word) - 1:len(word)] == "," {
                    word = word[:len(word) - 1]
                }
                children[word] = true
            }
        }
    }
    for _, program := range possibleRoot {
        if !children[program] {
            fmt.Println(program)
            break
        }
    }
}

