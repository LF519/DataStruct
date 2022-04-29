/*
图的邻接矩阵
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

// 建立无向网图的邻接矩阵表示
func CreateMGraph(m *MGraph) {
	var i, j, k, w int

	fmt.Println("输入顶点数和边数: ")
	fmt.Scan(&m.numVertexes, &m.numEdges) // 输入顶点数和边数

	// 初始化顶点表和邻接矩阵
	m.vexs = make([]string, m.numVertexes)
	m.arc = make([][]int, m.numVertexes)
	for i := range m.arc {
		m.arc[i] = make([]int, m.numVertexes)
		for j := range m.arc[i] {
			if i == j {
				m.arc[i][j] = 0
			} else {
				m.arc[i][j] = X
			}
		}
	}

	fmt.Println("请输入顶点信息: ")
	// 读入顶点信息, 建立顶点表
	for i = 0; i < m.numVertexes; i++ {
		fmt.Scan(&m.vexs[i])
	}

	// 读入numEdges条边, 建立邻接矩阵
	for k = 0; k < m.numEdges; k++ {
		fmt.Println("输入边(vi, vj)上的下标i, 下标j和权w: ")
		fmt.Scan(&i, &j, &w)
		m.arc[i][j] = w
		m.arc[j][i] = m.arc[i][j] // 因为是无向图, 矩阵对称
	}

}

// 邻接矩阵深度优先递归算法
func DFS(mg *MGraph, i int, v []bool) {
	v[i] = true
	fmt.Printf("mg.vexs[%d]: %v\n", i, mg.vexs[i]) // 打印顶点(也可以做其他操作)
	for j := 0; j < mg.numVertexes; j++ {

		// 对未访问的邻接顶点递归调用
		if mg.arc[i][j] > 0 && mg.arc[i][j] < X && !v[j] {
			DFS(mg, j, v)
		}
	}
}

// 邻接矩阵的深度遍历操作
func DFSTraverse(mg *MGraph) {
	vistied := make([]bool, mg.numVertexes) // 访问标记的切片
	// 对未访问过得顶点调用DFS
	// 如果是连通图, 只会执行一次(相当于不用for, 直接 DFS(mg, 0, vistied) )
	for i := 0; i < mg.numVertexes; i++ {
		if !vistied[i] {
			DFS(mg, i, vistied)
		}
	}
}

// 邻接矩阵广度优先递归算法
func BFSTraverse(mg *MGraph) {
	queue := make([]int, 0)                 // 初始化一个辅助用的队列
	vistied := make([]bool, mg.numVertexes) // 访问标记的切片
	// 对每一个顶点做循环, 如果是连通图, 只会执行一次
	for i := 0; i < mg.numVertexes; i++ {
		if !vistied[i] {
			vistied[i] = true                              // 设置当前顶点访问过
			fmt.Printf("mg.vexs[%d]: %v\n", i, mg.vexs[i]) // 打印顶点, 也可以做其他操作
			queue = append(queue, i)
			// 若当前队列不为空
			for len(queue) > 0 {
				i = queue[0] // 将队列中的元素取出, 赋值给i
				queue = queue[1:]
				for j := 0; j < mg.numVertexes; j++ {
					// 判断其他顶点若与当前顶点有边存在, 且从未被访问过
					if mg.arc[i][j] > 0 && mg.arc[i][j] < X && !vistied[j] {
						vistied[j] = true // 将找到的此顶点标记为已访问
						fmt.Printf("mg.vexs[%d]: %v\n", j, mg.vexs[j])
						queue = append(queue, j) // 将找到的此顶点加入队列
					}
				}
			}
		}
	}
}

func main() {
	// mg := new(MGraph)
	// CreateMGraph(mg)
	// fmt.Println(*mg)
	// DFSTraverse(mg)

	mg2 := MGraph{
		vexs: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"},
		arc: [][]int{
			//A B  C  D  E  F  G  H  I
			{0, 1, X, X, X, 1, X, X, X}, // A
			{1, 0, 1, X, X, X, 1, X, 1}, // B
			{X, 1, 0, 1, X, X, X, X, 1}, // C
			{X, X, 1, 0, 1, X, 1, 1, 1}, // D
			{X, X, X, 1, 0, 1, X, 1, X}, // E
			{1, X, X, X, 1, 0, 1, X, X}, // F
			{X, 1, X, 1, X, 1, 0, 1, X}, // G
			{X, X, X, 1, 1, X, 1, 0, X}, // H
			{X, 1, 1, 1, X, X, X, X, 0}, // I
		},
		numVertexes: 9,
		numEdges:    15,
	}
	fmt.Println(mg2)
	DFSTraverse(&mg2)
	fmt.Println("==============")
	BFSTraverse(&mg2)
}
