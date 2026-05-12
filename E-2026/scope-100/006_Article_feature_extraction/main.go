package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fragments := strings.Fields(scanner.Text())

	// 用数组的索引来当成key，val来作为计数('a'=10的,a-z就是从10开始 ch - 'a'就是对应的index)
	commonCount := make([]int, 26)
	for _, ch := range fragments[0] {
		commonCount[ch-'a']++
	}
	// 依次处理后续每个片段
	for i := 1; i < len(fragments); i++ {
		// 统计当前片段中每个字母的出现次数
		currentCount := make([]int, 26)
		for _, ch := range fragments[i] {
			currentCount[ch-'a']++
		}

		// 对每个字母更新为最小出现次数
		for j := 0; j < 26; j++ {
			if currentCount[j] < commonCount[j] {
				commonCount[j] = currentCount[j]
			}
		}
	}

	// 按字母从小到大拼接结果
	var result strings.Builder

	for i := 0; i < 26; i++ {
		for j := 0; j < commonCount[i]; j++ {
			result.WriteByte(byte('a' + i))
		}
	}

	// 输出最终特征提取结果
	fmt.Println(result.String())

}
