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
	jiaStrs := strings.Fields(scanner.Text())
	scanner.Scan()
	yiStrs := strings.Fields(scanner.Text())
	jias := make([]int, len(jiaStrs))
	yis := make([]int, len(yiStrs))
	for i := 0; i < len(jiaStrs); i++ {
		val, _ := strconv.Atoi(jiaStrs[i])
		jias[i] = val
	}
	for i := 0; i < len(yiStrs); i++ {
		val, _ := strconv.Atoi(yiStrs[i])
		yis[i] = val
	}

	tables := make([]int, 0)
	tableMap := make(map[int]int, 0)

	times := 0

	for {
		// 甲没牌
		if len(jias) <= 0 {
			if len(yis) > 0 {
				// 乙胜
				fmt.Println(yis[len(yis)-1])
			} else {
				if len(tables) <= 0 {
					fmt.Println(0)
				} else {
					fmt.Println(tables[len(tables)-1])
				}
			}
			return
		}
		// 甲出
		jias, tables = chupai(jias, tables, tableMap, &times)
		if times > 10000 {
			// 平局
			if len(tables) <= 0 {
				fmt.Println(0)
			} else {
				fmt.Println(tables[len(tables)-1])
			}
			return
		}
		// 乙没牌
		if len(yis) <= 0 {
			if len(jias) > 0 {
				// 甲胜
				fmt.Println(jias[len(jias)-1])
			} else {
				// 平局
				if len(tables) <= 0 {
					fmt.Println(0)
				} else {
					fmt.Println(tables[len(tables)-1])
				}
			}
			return
		}
		// 乙出
		yis, tables = chupai(yis, tables, tableMap, &times)
		if times > 10000 {
			// 平局
			if len(tables) <= 0 {
				fmt.Println(0)
			} else {
				fmt.Println(tables[len(tables)-1])
			}
			return
		}
	}

}

func chupai(hands []int, tables []int, tableMap map[int]int, times *int) ([]int, []int) {
	if *times > 10000 {
		return hands, tables
	}
	// 出牌
	chu := hands[0]
	hands = hands[1:]

	*times++

	if idx, ok := tableMap[chu]; ok {
		// 这张牌存在，收牌
		ret := tables[idx:]
		ret = append(ret, chu)
		// 翻面，放在底部
		for i := len(ret) - 1; i >= 0; i-- {
			delete(tableMap, ret[i])
			hands = append(hands, ret[i])
		}
		// 桌面剩余
		tables = tables[0:idx]
		// 继续出牌
		hands, tables = chupai(hands, tables, tableMap, times)
	} else if chu == 11 {
		// 打出的是J
		if len(tables) > 0 {
			// 桌面有牌，全部收走，翻面，放在底部，清空桌面
			for i := len(tables) - 1; i >= 0; i-- {
				hands = append(hands, tables[i])
			}
			tableMap = make(map[int]int)
			tables = []int{}
			// 继续出牌
			hands, tables = chupai(hands, tables, tableMap, times)
		} else {
			// 将J放在桌面
			tables = append(tables, chu)
		}
	} else {
		// 这张牌不存在,桌面加牌
		tables = append(tables, chu)
		// 记录出现的索引
		tableMap[chu] = len(tables) - 1
	}
	return hands, tables
}
