package nowcoder

type ListNode struct {
    Val  int
    Next *ListNode
}

// 重排链表
// {1,2,3,4} -> {1,4,2,3}
// {1,2,3,4,5} -> {1,5,2,4,3}
func ReorderList(head *ListNode) {}

/* 反转链表
思路：1 -> 2 -> 3 => 反转为 => 1 <- 2 <- 3，则仅需要遍历首位，将每一个节点 next 置为前一个即可
步骤：
1、初始化 pre 一个空指针，cur、next 两个指针，指向头节点
2、next 移动到下一位，cur 的下一个指向 pre
3、pre 指针指向 cur
4、cur 指针指向 next
5、遍历
*/
func ReverseList(head *ListNode) *ListNode {
    var (
        pre  *ListNode
        cur  = head
        next = head
    )
    for next != nil {
        next = cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre
}

/* 查找链表中间节点 - 偶数选右
思路：初始化两个指针，每次遍历，一个遍历一个单位，另一个遍历两个单位，直至不满足遍历
步骤：
1、初始化两个指针 cur、double
2、double 遍历两个单位，cur 遍历一个单位
3、直至 double 为空或者 double.Next 为空，此时 cur 为中间节点
*/
func MiddleRightList(head *ListNode) *ListNode {
    var (
        cur    = head
        double = head
    )
    for double != nil && double.Next != nil {
        double = double.Next.Next
        cur = cur.Next
    }
    return cur
}

/* 查找链表中间节点 - 偶数选左
思路：参考 MiddleRightList
*/
func MiddleLeftList(head *ListNode) *ListNode {
    var (
        cur    = head
        double = head
    )
    for double != nil && double.Next != nil {
        double = double.Next.Next
        if double == nil {
            break
        }
        cur = cur.Next
    }
    return cur
}
