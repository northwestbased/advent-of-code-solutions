package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)


func main() {
    scanner := bufio.NewScanner(os.Stdin)
    instructions := []int{}
    for scanner.Scan() {
        instruction := scanner.Text()
        integer, _ := strconv.Atoi(instruction)
        instructions = append(instructions, integer)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    } else {
        jumpcount := 0
        position := 0
        length := len(instructions)
        for position >= 0 && position < length {
            old_pos := position
            position += instructions[position]
            instructions[old_pos] += 1
            jumpcount += 1
        }
        fmt.Println(jumpcount)
    }
}
