package parse

import (
    "reflect"

    "mussl2/builtins"
)

func isCondOp(op string) bool {
    if op == "if" {
        return true
    } else {
        return false
    }
}

func CallOp(op string, opArgs []string) string {
    var res string
    switch op {
        case "+":
            res = builtins.Add(opArgs...)
        case "-":
            res = builtins.Subtract(opArgs...)
        case "==":
            res = builtins.Equals(opArgs...)
        case "print":
            res = builtins.Print(opArgs...)
        case "prog":
            res = builtins.Prog(opArgs...)
    }
    return res
}

func CallCondOp(op string, cond *List) string {
    var res string
    switch op {
        case "if":
            res = EvalExpression(cond) 
    }
    return res
}

func EvalExpression(ex *List) string {
    exOp := ex.head.value.(string)
    curNode := ex.head

    exArgs := make([]string, 0)
    if isCondOp(exOp) {
        cond := ex.head.next.value.(*List)
       
        var do *List
        if CallCondOp(exOp, cond) == "SUCCESS" {
            do = ex.head.next.next.value.(*List)
        } else {
            do = ex.head.next.next.next.value.(*List)
        }

        exArgs = append(exArgs, EvalExpression(do))
    } else {
        for curNode.next != nil {
            curNode = curNode.next

            if reflect.TypeOf(curNode.value) == reflect.TypeOf(ex) {
                exArgs = append(exArgs, EvalExpression(curNode.value.(*List)))
            } else {
                exArgs = append(exArgs, curNode.value.(string))
            }
        }
    }

    return CallOp(exOp, exArgs)
}
