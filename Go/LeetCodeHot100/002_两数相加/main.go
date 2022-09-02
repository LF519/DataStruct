/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例 1:
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]

示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, tmp, value1, value2 int

	currentNode := new(ListNode)
	headNode := currentNode // 记录链表最开头的地址

	for l1 != nil || l2 != nil {
		currentNode.Next = new(ListNode)
		currentNode = currentNode.Next
		if l1 != nil {
			value1 = l1.Val
			l1 = l1.Next
		} else {
			value1 = 0
		}

		if l2 != nil {
			value2 = l2.Val
			l2 = l2.Next
		} else {
			value2 = 0

		}

		sum = (value1 + value2 + tmp) % 10
		tmp = (value1 + value2 + tmp) / 10

		currentNode.Val = sum

	}

	if tmp > 0 {
		currentNode.Next = &ListNode{Val: tmp}
	}
	return headNode.Next // 因为循环里面一开头就给current添加了下一个节点, 所以headNode.Next才是要返回的头结点
}

func main() {
	l1 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	sumNode := addTwoNumbers(l1, l2)
	for sumNode != nil {
		fmt.Println(sumNode.Val)
		sumNode = sumNode.Next
	}
}
