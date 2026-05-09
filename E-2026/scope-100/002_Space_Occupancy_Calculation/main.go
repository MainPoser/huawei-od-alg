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
	scanner.Scan()
	target := strings.TrimSpace(scanner.Text())
	scanner.Scan()
	allPaths := strings.Fields(scanner.Text())
	scanner.Scan()
	allPathSizeStrs := strings.Fields(scanner.Text())
	allPathSizes := make([]int, len(allPathSizeStrs))
	for i, allPthSizeStr := range allPathSizeStrs {
		pathSize, _ := strconv.Atoi(allPthSizeStr)
		allPathSizes[i] = pathSize
	}

	subPathSize := make(map[string]int)
	max := 0
	for i, path := range allPaths {
		if strings.HasPrefix(path, target+"/") {
			sub := strings.Replace(path, target+"/", "", 1)
			depthOneSub := strings.SplitN(sub, "/", 2)[0]
			count := subPathSize[depthOneSub]
			subPathSize[depthOneSub] = count + allPathSizes[i]
			if subPathSize[depthOneSub] > max {
				max = subPathSize[depthOneSub]
			}
		}
	}
	var results []string
	for item, size := range subPathSize {
		if size == max {
			results = append(results, item)
		}
	}

	// 5. 排序并输出
	slices.Sort(results)
	fmt.Println(results[len(results)-1])
}
