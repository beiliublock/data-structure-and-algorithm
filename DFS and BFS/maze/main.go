package main

import "fmt"

func main()  {
	maze := readMaze("DFS and BFS/maze/maze.in")
	fmt.Println("迷宫地图为：")
	printMap(maze)
	//利用深度优先算法解决迷宫问题
	steps, b := DfsWalk(maze, point{0,0}, point{len(maze)-1,len(maze[0])-1})
	if b {
		fmt.Println("存在路径可以走出迷宫，行走路径为：")
	}else {
		fmt.Println("不存在路径可以走出迷宫，行走路径为：")
	}
	printMap(steps)
	step, b := BfsWalk(maze, point{0,0}, point{len(maze)-1,len(maze[0])-1})
	if b {
		fmt.Println("存在路径可以走出迷宫，行走路径为：")
	}else {
		fmt.Println("不存在路径可以走出迷宫，行走路径为：")
	}
	printMap(step)
}
