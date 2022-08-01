package parse

import (
    "reflect"
)

func isWhiteSpace(s string) bool {
    if s == " " || s == "\n" || s == "\t" {
        return true
    } else {
        return false
    }
}

func isReserved(s string) bool {
    res := []string{"(", ")"}
    for i := 0; i < len(res); i++ {
        if s == res[i] {
            return true
        }
    }
    return false
}

func climbToTopList(l *List) *List {
    curList := l
    for curList.upperList != nil {
        curList = curList.upperList
    }
    return l
}

func expressionToString(ex *List) string {
    exOp := ex.head.value.(string)
    curNode := ex.head

    exArgs := make([]string, 0)
    for curNode.next != nil {
        curNode = curNode.next

        if reflect.TypeOf(curNode.value) == reflect.TypeOf(ex) {
            exArgs = append(exArgs, expressionToString(curNode.value.(*List)))
        } else {
            exArgs = append(exArgs, curNode.value.(string))
        }
    }

    out := "(Operator: " + exOp + ", Args: "
    for i := 0; i < len(exArgs) - 1; i++ {
        out = out + exArgs[i] + ", "
    }
    out = out + exArgs[len(exArgs) - 1] + ")"
    return out
}
