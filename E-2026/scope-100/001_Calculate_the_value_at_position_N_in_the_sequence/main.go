package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		m, _ := strconv.Atoi(fields[0])
		n, _ := strconv.Atoi(fields[1])

		if n <= m {
			fmt.Println(n)
			continue
		}

		windows := make([]int, m)
		for i := 0; i < m; i++ {
			windows[i] = i + 1
		}
		max := slices.Max(windows)
		min := slices.Min(windows)
		val := max - min
		for i := m; i < n; i++ {
			fmt.Println(windows)
			if hasDup(windows) {
				val = max + min
			} else {
				val = max - min
			}
			windows = windows[1:]
			windows = append(windows, val)
			max = slices.Max(windows)
			min = slices.Min(windows)
		}
		fmt.Println(val)
	}
}

func hasDup(arr []int) bool {
	seen := make(map[int]struct{})
	for _, v := range arr {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}
