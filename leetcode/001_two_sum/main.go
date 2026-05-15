package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("两数之和")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("字母异位分组")
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println("最长连续序列")
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println("移动零")
	s1 := []int{0, 1, 0, 3, 12}
	moveZeroes(s1)
	fmt.Println(s1)
	fmt.Println("盛最多水的容器")
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println("三数之和")
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println("接雨水")
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println("无重复最长字符串")
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("找到字符串中所有字母异位词")
	fmt.Println(findAnagrams("cbaebabacd", "abc"))

	fmt.Println("买卖股票的最佳时机")
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	tmp := make(map[int]int)
	for i := range nums {
		if idx, ok := tmp[target-nums[i]]; ok {
			res[0] = idx
			res[1] = i
			break
		} else {
			tmp[nums[i]] = i
		}
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	tmp := make(map[string][]string)
	for i := range strs {
		strRune := []rune(strs[i])
		sort.Slice(strRune, func(i, j int) bool {
			return strRune[i] > strRune[j]
		})
		if _, ok := tmp[string(strRune)]; ok {
			tmp[string(strRune)] = append(tmp[string(strRune)], strs[i])
		} else {
			tmp[string(strRune)] = []string{strs[i]}
		}
	}
	res := make([][]string, 0)
	for word := range tmp {
		cur := make([]string, len(tmp[word]))
		copy(cur, tmp[word])
		res = append(res, cur)
	}
	return res
}

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	maxLentgth := 0
	tmp := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		} else if nums[i] == nums[i-1]+1 {
			tmp++
		} else {
			if tmp+1 > maxLentgth {
				maxLentgth = tmp + 1
			}
			tmp = 0
		}
	}
	if tmp+1 > maxLentgth {
		return tmp + 1
	}
	return maxLentgth
}

func moveZeroes(nums []int) {
	z := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[z], nums[i] = nums[i], nums[z]
			z++
		}
	}
}

func maxArea(height []int) int {
	max := 0
	l, r := 0, len(height)-1
	for l < r {
		if height[l] < height[r] {
			tmp := (r - l) * height[l]
			if tmp > max {
				max = tmp
			}
			l++
		} else {
			tmp := (r - l) * height[r]
			if tmp > max {
				max = tmp
			}
			r--
		}
	}

	return max
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		// --- 修改点 1：固定位去重 ---
		// 一定要和之前的比较，如果比较后面的n+1。会导致丢掉同时出现相同数字的答案
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			ans := nums[i] + nums[l] + nums[r]
			if ans == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// --- 修改点 2 & 3：左右指针去重 ---
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
				l++
			} else if ans < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// 接雨水
func trap(height []int) int {
	return 0
}

// 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	// 记录字符上一次出现的位置
	lastPos := make(map[rune]int)
	maxLen := 0
	left := 0 // 窗口左边界

	for i, char := range []rune(s) {
		// 如果当前字符在窗口内出现过
		if pos, ok := lastPos[char]; ok && pos >= left {
			// 左边界跳到重复字符位置的下一个
			left = pos + 1
		}

		// 更新字符最后出现的位置
		lastPos[char] = i

		// 计算当前窗口长度：右边界 - 左边界 + 1
		curLen := i - left + 1
		maxLen = max(curLen, maxLen)
	}

	return maxLen
}

// 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
    ns, np := len(s), len(p)
    if ns < np {
        return nil
    }

    res := []int{}
    // 使用数组代替 Map 统计字母频率（假设全是小写英文字母）
    var pCount, sCount [26]int
    for i := 0; i < np; i++ {
        pCount[p[i]-'a']++
        sCount[s[i]-'a']++
    }

    // 检查初始窗口
    if pCount == sCount {
        res = append(res, 0)
    }

    // 开始滑动
    for i := np; i < ns; i++ {
        // 右边进来一个
        sCount[s[i]-'a']++
        // 左边出去一个
        sCount[s[i-np]-'a']--

        // 比较两个数组是否相等（Go 支持数组直接比较）
        if sCount == pCount {
            res = append(res, i-np+1)
        }
    }

    return res
}

// 买卖股票的最佳时机
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// 初始化：第一天买入的价格就是最低价，利润为0
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// 如果今天的价格比之前的最低价还低，更新最低价
		minPrice = min(minPrice, prices[i])
		maxProfit = max(maxProfit, prices[i]-minPrice)
	}

	return maxProfit
}
