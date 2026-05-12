package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	lines := strings.Fields(scanner.Text())

	students := make([]string, 0)
	if err := json.Unmarshal([]byte(lines[0]), &students); err != nil {
		panic(err)
	}
	votes := make([]string, 0)
	if err := json.Unmarshal([]byte(lines[1]), &votes); err != nil {
		panic(err)
	}

	fmt.Println(students)
	fmt.Println(votes)

	if len(votes) > 3*len(students) {
		fmt.Println("Invalid election")
		return
	}

	stuVotes := make([]int, len(students))
	stuSegs := make([][]string, len(students))
	for i := range students {
		stuSegs[i] = strings.Split(students[i], "-")
	}
	for i := range votes {
		voteSegs := strings.Split(votes[i], "-")
		match := 0
		student := 0
		for j := range students {
			if matchStu(stuSegs[j], voteSegs) {
				match++
				student = j
			}
		}
		if match == 1 {
			stuVotes[student]++
		}
	}

	// 选举有效性检查
	if len(stuVotes) == 0 {
		fmt.Println("Invalid election")
		return
	}
	// 每个人的得票数不能大于n
	var results [][]int
	for idx, count := range stuVotes {
		// 单人票数不能超过总人数
		if count > len(students) {
			fmt.Println("Invalid election")
			return
		}
		results = append(results, []int{idx, count})
	}

	// 排序逻辑
	sort.Slice(results, func(i, j int) bool {
		if results[i][1] != results[j][1] {
			return results[i][1] > results[j][1] // 票数降序
		}
		return students[results[i][0]] < students[results[j][0]] // 名字升序
	})
	fmt.Println(students[results[0][0]])
}

func matchStu(stuSegs, voteSegs []string) bool {
	stuSegLen := len(stuSegs)
	voteSegLen := len(voteSegs)
	if voteSegLen > stuSegLen {
		return false
	}

	for i := 0; i <= stuSegLen-voteSegLen; i++ {
		match := true
		for j := 0; j < voteSegLen; j++ {
			if stuSegs[i+j] != voteSegs[j] {
				match = false
				break
			}
		}
		if match {
			return match
		}
	}

	return false

}
