package parse

import (
    "reflect"

    "mussl2/builtins"
)

func CallOp(op string, opArgs []string) string {
    var res string
    switch op {
        case "+":
            res = builtins.Add(opArgs...)
    }
    return res
}

func EvalExpression(ex *List) string {
    exOp := ex.head.value.(string)
    curNode := ex.head

    exArgs := make([]string, 0)
    for curNode.next != nil {
        curNode = curNode.next

        if reflect.TypeOf(curNode.value) == reflect.TypeOf(ex) {
            exArgs = append(exArgs, EvalExpression(curNode.value.(*List)))
        } else {
            exArgs = append(exArgs, curNode.value.(string))
        }
    }

    return CallOp(exOp, exArgs)
}
