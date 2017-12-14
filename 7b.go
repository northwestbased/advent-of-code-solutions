package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)


type program struct {
    children []string
    ownWeight int
    holdingWeight int
}


func findRootNode(programs map[string]program) string{
    childNodes := make(map[string]bool)
    for _, program := range programs {
        for _, c := range program.children {
            childNodes[c] = true
        }
    }
    for programName, _ := range programs {
        if !childNodes[programName] {
            return programName
        }
    }
    return ""
}


func setWeights(programName string) int {
    program := programs[programName]
    program.holdingWeight = program.ownWeight
    for _, c := range program.children {
        program.holdingWeight += setWeights(c)
    }
    programs[programName] = program
    return program.holdingWeight
}


func getBadBranch (programName string) string{

    // get program 
    program := programs[programName]



    // if the program has one child, the child subtree could be bad
    if len(program.children) == 1 {
        result :=  getBadBranch(program.children[0])
        if len(result) > 0 {
            return result
        }
        return ""


    // if there are two children, one of the subtrees is bad
    // otherwise, there could be more than one mistake
    } else if len(program.children) == 2 {
        result := getBadBranch(program.children[0])
        if len(result) > 0 {
            return result
        }
        result = getBadBranch(program.children[1])
        if len(result) > 0 {
            return result
        }
        return ""


    // otherwise, there are more than two children and one of them, or their trees, are bad
    } else {
        //if one of the children as a different weight than the others, return that program or one of the children
        for i := 0; i < len(program.children) - 2; i += 1 {
            w1 := programs[program.children[i]].holdingWeight
            w2 := programs[program.children[i+1]].holdingWeight
            w3 := programs[program.children[i+2]].holdingWeight
            if w1 == w2 && w1 != w3{
                result :=  getBadBranch(program.children[i+2])
                if len(result) > 0 {
                    return result
                }
                return programName
            } else if w1 == w3 && w1 != w2 {
                result := getBadBranch(program.children[i+1])
                if len(result) > 0 {
                    return result
                }
                return programName
            } else if w2 == w3 && w1 != w2{
                result := getBadBranch(program.children[i])
                if len(result) > 0 {
                    return result
                }
                return programName
            }
        }

        // if all of the children are the same weight, just iteratre through the subtrees - one of them is bad
        for _, p := range program.children {
            result := getBadBranch(p)
            if len(result) > 0 {
                return result
            }
        }
    }
    return ""
}


var programs map[string]program



func main() {

    programs = make(map[string]program)


    //get program input
    scanner := bufio.NewScanner(os.Stdin)
    var lines []string;

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }


    //put the programs into a map. we should reallyc reate a tree, but a map is simpler
    for _, line := range lines {
        splitline := strings.Split(line, " ")

        name := splitline[0]
        p := program{}

        p.ownWeight, _ = strconv.Atoi(splitline[1][1:len(splitline[1]) - 1]) //does this get the weight?

        if len(splitline) > 2 {
            for _, word := range splitline[3:] {
                if word[len(word) - 1:] == "," {
                    word = word[:len(word) - 1]
                }
                p.children = append(p.children, word)
            }
        }
        programs[name] = p
    }


    //find the root
    root := findRootNode(programs)
    setWeights(root)
    badbranch := getBadBranch(root)
    program := programs[badbranch]
    for i := 0; i <  len(program.children) - 2; i += 1 {
        w1 := programs[program.children[i]].holdingWeight
        w2 := programs[program.children[i+1]].holdingWeight
        w3 := programs[program.children[i+2]].holdingWeight
        if w1 == w2 && w1 != w3{
            fmt.Println(programs[program.children[i + 2]].ownWeight - (w3 - w2))
            break
        } else if w1 == w3 && w1 != w2 {
            fmt.Println(programs[program.children[i + 1]].ownWeight - (w2 - w3))
            break
        } else if w2 == w3 && w1 != w2{
            fmt.Println(programs[program.children[i]].ownWeight - (w1 - w2))
            break
        }
    }
}
