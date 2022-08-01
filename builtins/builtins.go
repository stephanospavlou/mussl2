package builtins

import "strconv"

func Add(nums ...string) string {
    total := 0

    for _, num := range nums { 
        rNum, _ := strconv.Atoi(num) 
        total = total + rNum 
    }

    return strconv.Itoa(total)
}
