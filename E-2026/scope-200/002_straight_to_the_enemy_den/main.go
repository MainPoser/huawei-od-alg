package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	sbIndexStrs := strings.Fields(scanner.Text())
	dx, _ := strconv.Atoi(sbIndexStrs[0])
	dy, _ := strconv.Atoi(sbIndexStrs[1])

	s := []int{0, n / 2}
	e := []int{n - 1, n / 2}

	fmt.Println(s)
	fmt.Println(e)
	fmt.Println(dx)
	fmt.Println(dy)
	gird := make([][]bool, n)
	for i := range n {
		gird[i] = make([]bool, n)
	}
	gird[s[0]][s[1]] = true
	count := 0
	minLength := math.MaxInt
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// 计算哨兵监视坐标
	dps := make([][]int, 0)
	dps = append(dps, []int{dx, dy})
	ddirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
	for _, ddir := range ddirs {
		ddx, ddy := dx+ddir[0], dy+ddir[1]
		if !(ddx >= n || ddy >= n || ddx < 0 || ddy < 0) {
			dps = append(dps, []int{ddx, ddy})
		}
	}

	dfs(gird, s[0], s[1], e[0], e[1], dps, dirs, &count, &minLength, 0)
	if count == 0 {
		fmt.Printf("%v %v", count, 0)
	} else {
		fmt.Printf("%v %v", count, minLength)
	}
}

func dfs(gird [][]bool, sx, sy, ex, ey int, dps [][]int, dirs [][]int, count *int, minLength *int, length int) {
	// 最小长度大于0，length已经大于最小长度，则直接返回
	if length > *minLength {
		return
	}
	// 走到了终点
	if sx == ex && sy == ey {
		if length < *minLength {
			*minLength = length
			*count = 1
		} else if length == *minLength {
			*count++
		}
	}

	// 向着4方向行走
	for _, dir := range dirs {
		nx, ny := sx+dir[0], sy+dir[1]
		// 判断是否越界，是否被哨兵发现
		if nx >= len(gird) || ny >= len(gird) || nx < 0 || ny < 0 {
			continue
		}
		match := false
		for _, dp := range dps {
			if dp[0] == nx && dp[1] == ny {
				match = true
				break
			}
		}
		if match {
			continue
		}
		// 如果已经走过，就不能走了
		if !gird[nx][ny] {
			gird[nx][ny] = true
			dfs(gird, nx, ny, ex, ey, dps, dirs, count, minLength, length+1)
			gird[nx][ny] = false
		}
	}
}
