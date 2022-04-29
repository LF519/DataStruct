/*
图的邻接矩阵
*/
package main

import (
	"fmt"
)

const INFINTY = 65535 // 用65535来表示无限大

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
				m.arc[i][j] = INFINTY
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
func main() {
	mg := new(MGraph)
	CreateMGraph(mg)
	fmt.Println(*mg)
}
