package main

import (
	"fmt"
)

//利用深度优先算法加回溯求第一条路径
//深度算法来解决迷宫问题的实现
//深度算法实现原理：
//取栈顶元素，求其邻接的未被访问的无障碍结点。求如果有，记其为已访问，并入栈。
//如果没有则回溯上一结点，具体做法是将当前栈顶元素出栈
//重复上述步骤，直到栈空（没有找到可行路径）或者栈顶元素等于终点（找到第一条可行路径）
func DfsWalk(grid [][]int, start, end point) ([][]int,bool) {
	steps := make([][]int, len(grid))
	for i:= range steps {
		steps[i] = make([]int, len(grid[i]))
	}
	//判断起始节点和终点是否合法
	_, t1 := start.at(grid)
	_, t2 := end.at(grid)
	if !(t1 && t2) {
		fmt.Println("输入的起始节点或者终点不合法")
		return steps,false
	}
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[len(Q) - 1]
		curlSteps, _ := cur.at(steps)
		curlen := len(Q)
		if cur == end {
			return steps,true
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			val, t := next.at(grid)
			if t && next != start{
				if val == 0 && steps[next.i][next.j] == 0 {
					steps[next.i][next.j] = curlSteps + 1
					Q = append(Q, next)
					break
				}
				//如果最后本轮最后一次还没有找到可行路径
				//则说明本轮没有可行路径，回退到上一个节点
				/*var s = point{1,0}
				if dir == s {
					Q = Q[:len(Q) - 1]
				}*/
			}
		}
		//当然也可以在本轮结束后检查是否找到可以路径
		//如果没有找到，则回退到上一个节点
		if len(Q) == curlen {
			Q = Q[:len(Q) - 1]
		}
	}
	return steps,false
}

