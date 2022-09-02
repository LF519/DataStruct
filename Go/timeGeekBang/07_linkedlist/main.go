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

	}
	return false
}

// MergeSortedList 两个有序链表合并
// 起一个新链表, 然后对l1和l2的当前节点进行比较, 谁的节点小, 就将该节点放到新链表里
// 然后该节点继续往下一位, 对l1和l2继续进行比较
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil || l1.head == nil || l1.head.next == nil {
		return l2
	}
	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return l1
	}

	newLinkList := &LinkedList{head: &ListNode{}}
	cur := newLinkList.head
	curl1 := l1.head.next
	curl2 := l2.head.next
	for curl1 != nil && curl2 != nil {
		if curl1.value.(int) > curl2.value.(int) {
			cur.next = curl2
			curl2 = curl2.next
		} else {
			cur.next = curl1
			curl1 = curl1.next
		}
		cur = cur.next

		if curl1 != nil {
			cur.next = curl1
		} else if curl2 != nil {
			cur.next = curl2
		}
	}

	return newLinkList
}

// DeleteBottomN 删除倒数第N个节点
// 使用快慢指针, 让快指针先走N步
func (ll *LinkedList) DeleteBottomN(n int) {
	if n <= 0 || ll.head == nil || ll.head.next == nil {
		return
	}

	fast := ll.head
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.next
	}

	if fast == nil {
		return
	}

	slow := ll.head
	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	slow.next = slow.next.next
}

// FindMiddleNode 获取中间节点
// 使用快慢指针
func (ll *LinkedList) FindMiddleNode() *ListNode {
	if ll.head == nil || ll.head.next == nil {
		return nil
	}
	if ll.head.next.next == nil {
		return ll.head.next
	}

	slow, fast := ll.head, ll.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
