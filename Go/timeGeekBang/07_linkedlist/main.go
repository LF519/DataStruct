package _7_linkedlist

import "fmt"

// ListNode 单链表节点
type ListNode struct {
	next  *ListNode
	value interface{}
}

// LinkedList 单链表(头结点)
type LinkedList struct {
	head *ListNode
}

// Print 打印链表
func (ll *LinkedList) Print() {
	cur := ll.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.value)
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

// Reverse 单链表翻转, 时间复杂度O(N)
// 反向链表思想：从前往后将每个节点的指针反向，即.next内的地址换成前一个节点(pre)
// 但为了防止后面链表的丢失, 在每次换之前需要先创建个指针指向下一个节点(tmp)
func (ll *LinkedList) Reverse() {
	if ll.head == nil || ll.head.next == nil || ll.head.next.next == nil {
		return
	}

	// 当时正在遍历的节点
	cur := ll.head.next
	// 正在遍历的节点的前一个节点
	//var pre *ListNode = nil
	pre := &ListNode{}
	for cur != nil {
		// 正在遍历的节点的后一个节点
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
}

// HasCycle 判断单链表是否有环
// 使用快慢指针
func (ll *LinkedList) HasCycle() bool {
	if ll.head != nil {
		slow := ll.head
		fast := ll.head
		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
		return false
	}
}

// MergeSortedList 两个有序链表合并
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil || l1.head == nil || l1.head.next == nil {
		return l2
	}
	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return l1
	}

	l = &LinkedList{head: &ListNode{}}

}
