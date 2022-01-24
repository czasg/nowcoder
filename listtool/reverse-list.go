package listtool

// 重排链表
// {1,2,3,4} -> {1,4,2,3}
// {1,2,3,4,5} -> {1,5,2,4,3}
func ReorderList(node *ListNode) {
    head := node
    count := 0
    for node.Next != nil {
        count++
        node = node.Next
    }
    tail := node
}
