/* 图的邻接表 */
package main

import "fmt"

// 边表节点
type EdgeNode struct {
	adjvex int       // 邻接点域, 存储顶点对应的下标
	weight int       // 用于存储权值, 对于非网图可以不需要
	next   *EdgeNode // 链域, 指向下一个邻接点
}

// 顶点表节点
type VertextNode struct {
	data      string    // 顶点域, 存储顶点信息
	firstedge *EdgeNode // 边表头指针
}

// 邻接表结构
type GraphAdjList struct {
	adjList     []VertextNode
	numVertexes int
	numEdges    int
}

// 建立图的邻接表结构
func CreateALGraph(g *GraphAdjList) {
	var i, j, k int
	var e *EdgeNode

	fmt.Println("输入顶点数和边数: ")
	fmt.Scan(&g.numVertexes, &g.numEdges) // 输入顶点数和边数

	// 读入顶点信息, 建立顶点表
	fmt.Println("输入顶点信息: ")
	for i = 0; i < g.numVertexes; i++ {
		g.adjList = append(g.adjList, VertextNode{})
		fmt.Scan(&g.adjList[i].data) // 输入顶点信息
	}

	// 建立边表, 应用的是单链表中的头插法
	for k = 0; k < g.numEdges; k++ {
		fmt.Println("输入边(vi, vj)上的顶点序号: ")
		fmt.Scan(&i, &j) // 输入vi, vj上的顶点序号

		e = new(EdgeNode)
		e.adjvex = j                    // 邻接序号为j
		e.next = g.adjList[i].firstedge // 将e指针指向当前顶点指向的结点
		g.adjList[i].firstedge = e      // 将当前顶点的指针指向e

		e = new(EdgeNode)
		e.adjvex = i                    // 邻接序号为i
		e.next = g.adjList[j].firstedge // 将e指针指向当前顶点指向的结点
		g.adjList[j].firstedge = e      // 将当前顶点的指针指向e
	}

}

// 邻接表的深度优先递归算法
func DFS(g *GraphAdjList, i int, v []bool) {
	var e *EdgeNode
	v[i] = true
	fmt.Printf("g.adjList[%d].data: %v\n", i, g.adjList[i].data) // 打印顶点, 也可以用于其他操作
	e = g.adjList[i].firstedge
	for e != nil {
		if !v[e.adjvex] {
			DFS(g, e.adjvex, v)
		}
		e = e.next
	}
}

// 邻接表的深度优先递归算法
func DFSTraverse(g *GraphAdjList) {
	vistied := make([]bool, g.numVertexes) // 初始化所有顶点的状态都是未访问过状态
	// 对未访问的顶点调用DFS, 若是连通图, 只会执行一次
	for i := 0; i < g.numVertexes; i++ {
		if !vistied[i] {
			DFS(g, i, vistied)
		}
	}
}

// 邻接表广度优先算法
func BFSTraverse(g *GraphAdjList) {
	queue := make([]int, 0)                // 初始化一个辅助用的队列
	vistied := make([]bool, g.numVertexes) // 访问标记的切片

	var e *EdgeNode
	for i := 0; i < g.numVertexes; i++ {
		if !vistied[i] {
			vistied[i] = true
			fmt.Printf("g.adjList[%d].data: %v\n", i, g.adjList[i].data) // 打印顶点, 也可以做其他操作
			queue = append(queue, i)
			// 若当前队列不为空
			for len(queue) > 0 {
				i = queue[0] // 将队列中的元素取出, 赋值给i
				queue = queue[1:]
				e = g.adjList[i].firstedge // 找到当前顶点边表的链表头指针
				for e != nil {
					// 若此顶点未被访问
					if !vistied[e.adjvex] {
						vistied[e.adjvex] = true                                                   // 将找到的此顶点标记为已访问
						fmt.Printf("g.adjList[%d].data: %v\n", e.adjvex, g.adjList[e.adjvex].data) // 打印顶点, 也可以做其他操作
						queue = append(queue, e.adjvex)                                            // 将此顶点入队列
					}
					e = e.next // 指针指向下一个邻接点
				}
			}
		}
	}
}

func main() {
	// gl := new(GraphAdjList)
	// CreateALGraph(gl)
	// fmt.Println(gl)
	g2 := GraphAdjList{
		// 下标对对应关系: A0 B1 C2 D3 E4 F5 G6 H7 I8
		adjList: []VertextNode{
			{data: "A", firstedge: &EdgeNode{adjvex: 1, weight: 1, next: &EdgeNode{adjvex: 5, weight: 1}}},                                                                                                                      // B F
			{data: "B", firstedge: &EdgeNode{adjvex: 0, weight: 1, next: &EdgeNode{adjvex: 2, weight: 1, next: &EdgeNode{adjvex: 6, weight: 1, next: &EdgeNode{adjvex: 8, weight: 1}}}}},                                        // A C G I
			{data: "C", firstedge: &EdgeNode{adjvex: 1, weight: 1, next: &EdgeNode{adjvex: 3, weight: 1, next: &EdgeNode{adjvex: 8, weight: 1}}}},                                                                               // B D I
			{data: "D", firstedge: &EdgeNode{adjvex: 2, weight: 1, next: &EdgeNode{adjvex: 4, weight: 1, next: &EdgeNode{adjvex: 6, weight: 1, next: &EdgeNode{adjvex: 7, weight: 1, next: &EdgeNode{adjvex: 8, weight: 1}}}}}}, // C E G H I
			{data: "E", firstedge: &EdgeNode{adjvex: 3, weight: 1, next: &EdgeNode{adjvex: 5, weight: 1, next: &EdgeNode{adjvex: 7, weight: 1}}}},                                                                               // D F H
			{data: "F", firstedge: &EdgeNode{adjvex: 0, weight: 1, next: &EdgeNode{adjvex: 4, weight: 1, next: &EdgeNode{adjvex: 6, weight: 1}}}},                                                                               // A E G
			{data: "G", firstedge: &EdgeNode{adjvex: 1, weight: 1, next: &EdgeNode{adjvex: 3, weight: 1, next: &EdgeNode{adjvex: 5, weight: 1, next: &EdgeNode{adjvex: 7, weight: 1}}}}},                                        // B D F H
			{data: "H", firstedge: &EdgeNode{adjvex: 3, weight: 1, next: &EdgeNode{adjvex: 4, weight: 1, next: &EdgeNode{adjvex: 6, weight: 1}}}},                                                                               // D E G
			{data: "I", firstedge: &EdgeNode{adjvex: 1, weight: 1, next: &EdgeNode{adjvex: 2, weight: 1, next: &EdgeNode{adjvex: 3, weight: 1}}}},                                                                               // B C D
		},
		numVertexes: 9,
		numEdges:    15,
	}

	var e *EdgeNode
	for _, v := range g2.adjList {

		fmt.Println(v.data)
		e = v.firstedge
		for e != nil {
			fmt.Print(g2.adjList[e.adjvex])
			e = e.next
		}
		fmt.Println()
	}
	DFSTraverse(&g2)
	fmt.Println("=============================")
	BFSTraverse(&g2)
}
