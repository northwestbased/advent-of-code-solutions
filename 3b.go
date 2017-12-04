package main


import (
    "fmt"
    //"bufio"
    //"os"
    "strconv"
)

const (
    LEFT = 0
    RIGHT = 1
    UP = 2
    DOWN = 3
)

const input int = 368078


func main() {
    values := make(map[string]int)
    direction := RIGHT

    values["0,0"] = 1;
    values["1,0"] = 1;
    posX := 1
    posY := 0
    currentvalue := 0

    getVal := func (x int, y int) int {
        key := strconv.Itoa(x) + "," + strconv.Itoa(y)
        return values[key]
    }
    setVal := func (x int, y int, value int) {
        key := strconv.Itoa(x) + "," + strconv.Itoa(y)
        values[key] = value
    }

    for input >= currentvalue{

        if direction == RIGHT && getVal(posX, posY + 1) == 0 {
            direction = UP
        } else if direction == UP && getVal(posX - 1, posY) == 0 {
            direction = LEFT
        } else if direction == LEFT && getVal(posX, posY - 1) == 0 {
            direction = DOWN
        } else if direction == DOWN && getVal(posX + 1, posY) == 0 {
            direction = RIGHT
        }

        if direction == RIGHT {
            posX++
        } else if direction == DOWN {
            posY--
        } else if direction == UP {
            posY++
        } else if direction == LEFT {
            posX--
        }


        currentvalue = getVal(posX + 1, posY) + getVal(posX + 1, posY + 1) + getVal(posX + 1, posY - 1)
        currentvalue += getVal(posX - 1, posY) + getVal(posX - 1, posY + 1) + getVal(posX - 1, posY - 1)
        currentvalue += getVal(posX, posY + 1) + getVal(posX, posY - 1)
        setVal(posX, posY, currentvalue)
    }
    fmt.Println(currentvalue)
}


