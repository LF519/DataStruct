package main

import "fmt"

/* 二叉链表结点结构定义代码 */
type BitNode struct {
	Data   string   // 结点数据
	LChild *BitNode //左指针
	RChild *BitNode // 右指针
}

// 1. 前序遍历, 从根结点就开始遍历左子树, 遍历到左边的尽头时, 找上一个结点的右子结点
// 2. 中序遍历, 先从根结点开始, 找到最左边的尽头开始遍历, 然后往上一级, 再往下一级右结点
// 3. 后序遍历, 先从根结点开始, 找到最左边的尽头开始遍历, 然后先左后右, 再往上一级
// 4. 层序遍历, 从根结点开始遍历, 先从上到下, 再从左到右

// 前序遍历递归算法
func PreOrderTraverse(tree *BitNode) {
	if tree == nil {
		return
	}
	fmt.Printf("%v", tree.Data)
	PreOrderTraverse(tree.LChild)
	PreOrderTraverse(tree.RChild)
}

// 中序遍历递归算法
func InOrderTraverse(tree *BitNode) {
	if tree == nil {
		return
	}
	InOrderTraverse(tree.LChild)
	fmt.Printf("%v", tree.Data)
	InOrderTraverse(tree.RChild)
}

// 后序遍历递归算法
func PostOrderTraverse(tree *BitNode) {
	if tree == nil {
		return
	}
	PostOrderTraverse(tree.LChild)
	PostOrderTraverse(tree.RChild)
	fmt.Printf("%v", tree.Data)
}

func main() {
	bitTree := BitNode{
		Data: "A",
		LChild: &BitNode{
			Data: "B",
			LChild: &BitNode{
				Data: "D",
				LChild: &BitNode{
					Data: "H",
					RChild: &BitNode{
						Data: "K",
					},
				},
			},
			RChild: &BitNode{
				Data: "E",
			},
		},
		RChild: &BitNode{
			Data: "C",
			LChild: &BitNode{
				Data: "F",
				LChild: &BitNode{
					Data: "I",
				},
			},
			RChild: &BitNode{
				Data: "G",
				RChild: &BitNode{
					Data: "J",
				},
			},
		},
	}
	PreOrderTraverse(&bitTree)
	fmt.Println()
	InOrderTraverse(&bitTree)
	fmt.Println()
	PostOrderTraverse(&bitTree)
}
