package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	s1 := []int{0, 1, 0, 3, 12}
	moveZeroes(s1)
	fmt.Println(s1)

	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

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
