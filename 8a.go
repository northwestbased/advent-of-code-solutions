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

    commands := make([][]string, 0)
    values := make(map[string]int)

    for scanner.Scan() {
        row := strings.Split(scanner.Text(), " ")
        commands = append(commands, row)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)

    } else {
        for _, c := range commands {
            register := c[0]
            operator := c[1]
            value, _ := strconv.Atoi(c[2])

            conditional_register := c[4]
            conditional_operator := c[5]
            conditional_operator_value := values[conditional_register]
            conditional_value, _ := strconv.Atoi(c[6])

            flag := false
            if conditional_operator == "!=" && conditional_operator_value != conditional_value {
                flag = true
            } else if conditional_operator == "==" && conditional_operator_value == conditional_value {
                flag = true
            } else if conditional_operator == ">" && conditional_operator_value > conditional_value {
                flag = true
            } else if conditional_operator == ">=" && conditional_operator_value >= conditional_value {
                flag = true
            } else if conditional_operator == "<" && conditional_operator_value < conditional_value {
                flag = true
            } else if conditional_operator == "<=" && conditional_operator_value <= conditional_value {
                flag = true
            }
            if flag {
                if operator == "inc" {
                    values[register] += value
                } else if operator == "dec" {
                    values[register] -= value
                }
            }
        }
    }
    max_val := values[commands[0][0]]
    for _, val := range values {
        if val > max_val {
            max_val = val
        }
    }
    fmt.Println(max_val)
}
