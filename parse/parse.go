package parse

import (
    "bufio"
    "fmt"
    "os"
)

func Tokenize(file string) []string {
    fd, err := os.Open(file)  
    if err != nil {
        fmt.Println("Error opening file. Maybe it doesn't exist?")
        os.Exit(-1) 
    }
    
    defer fd.Close()

    s := bufio.NewScanner(fd)
    s.Split(bufio.ScanRunes)

    tok := make([]string, 0)

    word := ""
    for s.Scan() {
        nextChar := s.Text()

        if !isWhiteSpace(nextChar) && !isReserved(nextChar) {
            word = word + nextChar
        } else {
            if word != "" {
                tok = append(tok, word)
                word = ""
            }
            if !isWhiteSpace(nextChar) {
                tok = append(tok, nextChar)
            }
        }
    }

    return tok
}

func Parse(file string) {
    tok := Tokenize(file)

    var curList *List
    for i := 0; i < len(tok); i++ {
        switch tok[i] {
            case "(":
                curList = curList.NewList()
            case ")":
                curList = curList.CloseList()
            default:
                curList.NewNode(tok[i])
        }
    }

    for i := 0; i < len(tok); i++ {
        fmt.Println(tok[i])
    }

    topList := climbToTopList(curList)
    fmt.Println(expressionToString(topList))

    fmt.Println("Result of script interpretation: " + EvalExpression(topList))
}
