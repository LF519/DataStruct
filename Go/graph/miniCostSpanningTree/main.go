/*
最小生成树
以邻接矩阵为例
*/
package main

import (
	"fmt"
)

const X = 65535 // 用65535来表示无限大

// 定义一个邻接矩阵图
type MGraph struct {
	vexs        []string // 顶点表
	arc         [][]int  // 邻接接矩阵, 可以看做边表
	numVertexes int      // 图中当前的顶点数
	numEdges    int      // 图中当前的边数
}

// 普利姆(Prim)算法生成最小生成树
func MiniSpanTreePrim(mg *MGraph) {
	// var min, k int

	adjvex := make([]int, mg.numVertexes)  // 保存相关顶点下标
	lowcost := make([]int, mg.numVertexes) // 保存相关顶点间边的权值

	lowcost[0] = 0 // 初始化第一个权值为0, 即v0加入生成树(lowcast的值为0, 在这里就是此下标的顶点已加入生成树)
	adjvex[0] = 0  // 初始化第一个顶点下标为0
	// 循环除下标为0外的全部顶点
	for i := 1; i < mg.numVertexes; i++ {
		lowcost[i] = mg.arc[0][i] // 将v0顶点与之有边的权值存入数组
		adjvex[i] = 0             // 初始化都为v0的下标
	}

	// i其实只管循环次数, 每次计算权值实际上是从adjvex中新加入的顶点开始
	// lowcost记录的是已经加入的顶点组成的总体与到达剩下其他顶点的权值数值
	// 每次新加入顶点后, lowcost都会发生相应的变化
	var min, k int
	for i := 1; i < mg.numVertexes; i++ {
		min = X // 初始化最小权值为无限大
		k = 0
		for j := 1; j < mg.numVertexes; j++ {
			// 如果权值不为0且小于min
			if lowcost[j] != 0 && lowcost[j] < min {
				min = lowcost[j] // 则让当前权值成为最小值
				k = j            // 记录最小值的下标为k
			}
		}

		fmt.Printf("(%d, %d)\n", adjvex[k], k) // 打印当前顶点边中权值最小边 (整体中到k最小的顶点, k)
		lowcost[k] = 0                         // 将当前顶点的权值设置为0, 表示此顶点已经完成任务

		// 循环所有顶点, 修改lowcost的值
		// adjvex数据描述的是当已加入顶点视作一个总体时, 这个总体中到达剩下的各(adjvex下标)顶点最近的顶点
		// 若lowcost数组为[0 0 18 65535 26 0 16 65535 12], adjvex的值为[0 0 1 0 5 0 1 0 1]时
		// 表示v0,v1,v5(lowcost值为0的下标)看做一个整体, 该整体到达v2的距离最近的顶点是整体的中的v1, 到达v4的距离最近的顶点是整体中的v5
		for j := 1; j < mg.numVertexes; j++ {
			if lowcost[j] != 0 && mg.arc[k][j] < lowcost[j] {
				lowcost[j] = mg.arc[k][j] // 将较小的权值存入lowcast
				adjvex[j] = k             // 将下标为k的顶点存入adjvex
			}
		}
	}
	fmt.Println()
	fmt.Printf("adjvex: %v\n", adjvex)
	fmt.Printf("lowcost: %v\n", lowcost)
}

func main() {
	mg := MGraph{
		vexs: []string{"v0", "v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8"},
		arc: [][]int{
			{0, 10, X, X, X, 11, X, X, X},   // v0
			{10, 0, 18, X, X, X, 16, X, 12}, // v1
			{X, X, 0, 22, X, X, X, X, 8},    // v2
			{X, X, 22, 0, 20, X, X, 16, 21}, // v3
			{X, X, X, 20, 0, 26, X, 7, X},   // v4
			{11, X, X, X, 26, 0, 17, X, X},  // v5
			{X, 16, X, X, X, 17, 0, 19, X},  // v6
			{X, X, X, 16, 7, X, 19, 0, X},   // v7
			{X, 12, 8, 21, X, X, X, X, 0},   // v8
		},
		numVertexes: 9,
		numEdges:    15,
	}

	MiniSpanTreePrim(&mg)
}
