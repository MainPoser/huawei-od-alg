package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	args := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(args[0])
	m, _ := strconv.Atoi(args[1])
	step, _ := strconv.Atoi(args[2])

	maxPos := []int{0, 0}
	minPos := []int{0, 0}
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		xDataStrs := strings.Fields(scanner.Text())
		xDatas := make([]int, m)
		grid[i] = xDatas
		for j, xDataStr := range xDataStrs {
			xData, _ := strconv.Atoi(xDataStr)
			grid[i][j] = xData
			if xData >= grid[maxPos[0]][maxPos[1]] {
				maxPos[0] = i
				maxPos[1] = j
			}
			if xData <= grid[minPos[0]][minPos[1]] {
				minPos[0] = i
				minPos[1] = j
			}
		}
	}

	fmt.Println(maxPos)
	fmt.Println(minPos)
	fmt.Println(step)

	// 路径计数
	count := 0
	// 访问标记
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	// 定义方向
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited[minPos[0]][minPos[1]] = true
	dfs(minPos[0], minPos[1], maxPos[0], maxPos[1], &count, dirs, grid, visited, n, m, step)
	fmt.Println(count)
}

func dfs(sx, sy, ex, ey int, count *int, dirs [][]int, grid [][]int, visited [][]bool, n, m int, step int) {
	// 到达最高点（终点）
	if sx == ex && sy == ey {
		*count++
		return
	}

	for _, d := range dirs {
		nx, ny := sx+d[0], sy+d[1]

		// 1. 越界检查
		if nx < 0 || ny < 0 || nx >= n || ny >= m {
			continue
		}

		// 2. 访问检查 + 递增约束 + 高度差约束
		diff := grid[nx][ny] - grid[sx][sy]
		if !visited[nx][ny] && diff > 0 && diff <= step {
			visited[nx][ny] = true
			dfs(nx, ny, ex, ey, count, dirs, grid, visited, n, m, step)
			visited[nx][ny] = false // 回溯
		}
	}
}
