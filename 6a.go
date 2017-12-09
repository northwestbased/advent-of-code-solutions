package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func getLargestIndex(slice []int) int {
    highest := slice[0]
    highest_index := 0
    for index, val := range(slice) {
        if val > highest {
            highest_index = index
            highest = val
        }
    }
    return highest_index
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    banks_s := []int{}
    bank := []string{}
    for scanner.Scan() {
        bank = strings.Split(scanner.Text(), "\t")
    }

    for _, b := range bank {
        bank_int, _ := strconv.Atoi(b)
        banks_s = append(banks_s, bank_int)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)

    } else {
        length := len(banks_s)
        banks := make([]int, length)
        copy(banks, banks_s)

        rounds := 0


        seen := make(map[string]bool)

        for true {
            index := getLargestIndex(banks)
            count := banks[index]
            banks[index] = 0

            for count > 0 {
                count -= 1
                index += 1
                banks[index % length] += 1
            }
            rounds += 1
            key := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(banks)), ","), "[]")
            if seen[key] {
                break
            }
            seen[key] = true
        }
        fmt.Println(rounds)
    }
}
