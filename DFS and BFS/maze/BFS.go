package main

import "fmt"

//利用广度优先算法求迷宫的最短路径
//从入口节点开始，判断它上下左右的临节点是否满足条件，如果满足就入队列
//取队首元素并出队列，寻其未被访问的元素，将其入队列并标记元素的前驱节点为队首元素
//重复上述步骤，直到栈空或者找到了终点
func BfsWalk(grid [][]int, start, end point) ([][]int, bool) {
	steps := make([][]int, len(grid))
	for i := range steps {
		steps[i] = make([]int, len(grid[i]))
	}
	_, t1 := start.at(grid)
	_, t2 := end.at(grid)
	if !(t1 && t2) {
		fmt.Println("输入的起始节点或者终点不合法")
		return steps,false
	}
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		if cur == end {
			return steps,true
		}
		curSteps, _ := cur.at(steps)
		Q = Q[1:]
		for _,dir := range dirs {
			next := cur.add(dir)
			val, b := next.at(grid)
			if b && next != start {
				if val == 0 && steps[next.i][next.j] == 0 {
					steps[next.i][next.j] = curSteps + 1
					Q = append(Q, next)
				}
			}
		}
	}
	return steps,false
}


