package main

import (
	"fmt"
	"os"
)

//迷宫问题中用到的工具集

//将迷宫文件转换为二维数组
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	//读取文件第一行的迷宫行列
	var row, col int
	fmt.Fscanf(file,"%d %d",&row, &col)
	//通过读取到的行列来创建二维数组读取迷宫
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file,"%d",&maze[i][j])
		}
	}
	return maze
}

//定义节点结构体
type point struct {
	//注意这里为了避免混淆是按照数组来定义坐标的
	//和我们日常生活中的x,y相反
	i, j int
}

//声明前进的4种方式
var dirs = [4]point {
	{0,-1},//左
	{1,0},//下
	{0,1},//右
	{-1,0},//上
}

//实现节点前进的方法,这里我们为了不改变原节点的值所有不用指针
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

//返回当前节点的值(用于判断是否是障碍和是否走过)和当前节点是否合法
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i > (len(grid) - 1) {
		return 0,false
	}
	if p.j < 0 || p.j > (len(grid[p.i]) - 1) {
		return 0,false
	}
	return grid[p.i][p.j],true
}

//打印二维数组
func printMap(maze [][]int) {
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%3d",val)
		}
		fmt.Println()
	}
}
