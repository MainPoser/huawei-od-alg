package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	counts := make(map[int]int)
	n := len(line)
	for i := 0; i < n; {
		char := line[i]
		var key int
		// 判定逻辑：优先匹配双字符转义
		if char == 'u' && i+1 < n && line[i+1] == 'u' {
			key = getKeyVal('j') // uu 转义为 j
			i += 2
		} else if char == 't' && i+1 < n && line[i+1] == 't' {
			key = getKeyVal('b') // tt 转义为 b
			i += 2
		} else {
			key = getKeyVal(char) // 正常按键
			i += 1
		}
		counts[key]++
	}

	// 将 map 转换为二维切片以便排序
	var result [][]int
	for k, v := range counts {
		result = append(result, []int{k, v})
	}
	// 排序逻辑：
	// 1. 次数降序 (x[1])
	// 2. 转义值升序 (x[0])
	sort.Slice(result, func(i, j int) bool {
		if result[i][1] != result[j][1] {
			return result[i][1] > result[j][1]
		}
		return result[i][0] < result[j][0]
	})

	// 输出结果格式为 [[key, count], ...]
	fmt.Print("[")
	for i, res := range result {
		fmt.Printf("[%d, %d]", res[0], res[1])
		if i < len(result)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func getKeyVal(c byte) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}
	return int(c-'a') + 10
}
