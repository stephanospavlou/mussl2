package builtins

import (
    "fmt"
    "strconv"
)

func Add(nums ...string) string {
    total := 0

    for _, num := range nums { 
        rNum, _ := strconv.Atoi(num) 
        total = total + rNum 
    }

    return strconv.Itoa(total)
}

func Subtract(nums ...string) string {
    total, _ := strconv.Atoi(nums[0])
    
    for i := 1; i < len(nums); i++ {
        rNum, _ := strconv.Atoi(nums[i])
        total = total - rNum
    }

    return strconv.Itoa(total)
}

func Equals(args ...string) string {
    arg1, arg2 := args[0], args[1]
    if arg1 == arg2 {
        return "SUCCESS"
    } else {
        return "FAILURE"
    }
}

func Print(outs ...string) string {
    s := ""
    for _, out := range outs {
        s = s + " " + out
    }
    fmt.Println(s)
    return "SUCCESS"
}

func Prog(args ...string) string {
    for _, arg := range args {
        if arg == "FAILURE" {
            return "FAILURE"
        }
    }
    return "SUCCESS"
}
