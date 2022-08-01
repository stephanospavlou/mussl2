package main

import (
    "fmt"
    "os"

    "mussl2/parse"
)

func main() {
    const usageMsg = "mussl2 is the Mussl Lisp dialect interpreter" +
                        "\n    usage: mussl <file.mu>\n"

    if len(os.Args) != 2 {
        fmt.Println(usageMsg)
        return
    } else {
        fmt.Println("Parsing file: " + os.Args[1])
        parse.Parse(os.Args[1])
    }
}
