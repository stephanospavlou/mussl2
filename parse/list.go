package parse

type Node struct {
    value interface{}
    next *Node
}

type List struct {
    head *Node
    upperList *List

    curNode *Node
}

func (curList *List) NewNode(value interface{}) {
    n := new(Node)
    n.value = value
    n.next = nil

    if curList.curNode != nil {
        curList.curNode.next = n
    } else {
        curList.head = n
    }
    curList.curNode = n
}

func (curList *List) NewList() *List {
    l := new(List)
    l.head = nil
    l.upperList = curList
    l.curNode = nil

    if curList != nil {
        curList.NewNode(l)
    }
    return l
}

func (curList *List) CloseList() *List {
    if curList.upperList != nil {
        return curList.upperList
    } else {
        return curList
    }
}
