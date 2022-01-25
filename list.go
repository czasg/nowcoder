package nowcoder

type ListNode struct {
    Val  int
    Next *ListNode
}

/* 重排链表 - {1,2,3,4} -> {1,4,2,3}
思路：从左右两端开始，各取一个重新组织链表
步骤：
1、先找到链表中间节点
2、将中间节点及其后的链表反转
3、遍历 head 至中间节点，各取一个节点组织新链表
*/
func ReorderList(head *ListNode) *ListNode {
    left := head
    right := ReverseList(MiddleLeftList(head))
    for left != nil {
        leftNext := left.Next
        rightNext := right.Next
        left.Next = right
        right.Next = leftNext
        left = leftNext
        right = rightNext
    }
    return head
}

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

/* 查找链表中间节点 - 偶数选右 - 不破坏原链表
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

/* 查找链表中间节点 - 偶数选左 - 不破坏原链表
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
