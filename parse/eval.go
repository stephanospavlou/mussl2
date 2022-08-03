package parse

import (
    "reflect"

    "mussl2/builtins"
)

var sets = make(map[string]string)
var defs = make(map[string]*List)

func isCondOp(op string) bool {
    if op == "if" {
        return true
    } else {
        return false
    }
}

func isDeclareOp(op string) bool {
    if op == "set" || op == "def" {
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
        case "*":
            res = builtins.Multiply(opArgs...)
        case "==":
            res = builtins.Equals(opArgs...)
        case "!=":
            res = builtins.NotEquals(opArgs...)
        case "print":
            res = builtins.Print(opArgs...)
        case "get":
            res = builtins.Get()
        case "list":
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
    if exOp == "call" {
        return EvalExpression(defs[ex.head.next.value.(string)])
    } else if isDeclareOp(exOp) {
        if exOp == "def" {
            defs[ex.head.next.value.(string)] = ex.head.next.next.value.(*List)
        } else if exOp == "set" {
            if reflect.TypeOf(ex.head.next.next.value) == reflect.TypeOf(ex) {
                sets[ex.head.next.value.(string)] = EvalExpression(ex.head.next.next.value.(*List))
            } else {
                sets[ex.head.next.value.(string)] = ex.head.next.next.value.(string)
            }
        }
        return "SUCCESS"
    } else if isCondOp(exOp) {
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
                if l, found := defs[curNode.value.(string)]; found {
                    exArgs = append(exArgs, EvalExpression(l))
                } else {
                    exArgs = append(exArgs, curNode.value.(string))
                }
            }
        }
    }

    // check args for set vars
    for i := 0; i < len(exArgs); i++ {
        if s, found := sets[exArgs[i]]; found {
            exArgs[i] = s
        }
    }

    return CallOp(exOp, exArgs)
}
