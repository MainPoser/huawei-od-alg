package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("两数之和")
	// fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

	fmt.Println("字母异位词分组")
	// fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// fmt.Println(groupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

	fmt.Println("最长连续序列")
	// fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	// fmt.Println(longestConsecutive2([]int{100, 4, 200, 1, 3, 2}))

	fmt.Println("移动零")
	moveZeroData := []int{0, 1, 0, 3, 12}
	// moveZeroes(moveZeroData)
	fmt.Println(moveZeroData)

	fmt.Println("盛最多水的容器")
	// fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	fmt.Println("三数之和")
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))

	fmt.Println("无重复字符的最长子串")
	// fmt.Println(lengthOfLongestSubstring("abcabcbb"))

	fmt.Println("找到字符串中所有字母异位词")
	// fmt.Println(findAnagrams("cbaebabacd", "abc"))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type DoubleListNode struct {
	Val  int
	Prev *DoubleListNode
	Next *DoubleListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 两数之和
func twoSum(nums []int, target int) []int {
	exists := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if idx, ok := exists[target-nums[i]]; ok {
			return []int{idx, i}
		} else {
			exists[nums[i]] = i
		}
	}
	return []int{0, 0}
}

// 所有异位词分组
func groupAnagrams(strs []string) [][]string {

	groupMap := make(map[[26]byte][]string)

	for i := range strs {
		key := [26]byte{}
		for j := range strs[i] {
			key[strs[i][j]-'a']++
		}
		groupMap[key] = append(groupMap[key], strs[i])

	}

	groupSlice := make([][]string, 0)
	for _, val := range groupMap {
		groupSlice = append(groupSlice, val)
	}

	return groupSlice

}

// 最长连续序列
func longestConsecutive(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	sort.Ints(nums)
	ans := 1
	tmp := 1

	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			continue
		} else if nums[i] == nums[i-1]+1 {
			tmp++

		} else {
			tmp = 1
		}
		ans = max(tmp, ans)
	}
	return ans
}

// 最长连续序列 版本2 无需排序
func longestConsecutive2(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0

	for num := range numSet {
		// 当前数字的-1不存在，则是起点
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			longestStreak = max(longestStreak, currentStreak)
		}
	}
	return longestStreak
}

// 移动零
func moveZeroes(nums []int) {
	l := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}
}

// 盛最多水的容器
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		x := r - l
		h := min(height[l], height[r])
		ans = max(x*h, ans)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func maxArea2(height []int) int {
	l := 0
	r := len(height) - 1
	res := 0
	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res

}

// 三数之和 TODO 切记去重 锚点i和双指针l,r都需要
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			tmpSum := nums[i] + nums[l] + nums[r]
			if tmpSum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
				l++
			} else if tmpSum < 0 {
				l++
			} else {
				r--
			}
		}

	}

	return ans
}

// 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	exists := make(map[byte]int)
	l := 0

	for i := range s {
		if idx, ok := exists[s[i]]; ok && idx >= l {
			l = idx + 1
		}
		exists[s[i]] = i
		maxLen = max(i-l+1, maxLen)
	}

	return maxLen
}

// 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	ns, np := len(s), len(p)
	if ns < np {
		return []int{}
	}
	sCount, pCount := [26]int{}, [26]int{}

	for i := range p {
		sCount[s[i]-'a']++
		pCount[p[i]-'a']++
	}
	ans := make([]int, 0)
	if sCount == pCount {
		ans = append(ans, 0)
	}
	// 滑动
	for i := np; i < ns; i++ {
		sCount[s[i]-'a']++
		sCount[s[i-np]-'a']--
		if sCount == pCount {
			ans = append(ans, i-np+1)
		}

	}
	return ans
}

// 和为k的子数组
func subarraySum(nums []int, k int) int {
	sumCount := make(map[int]int)
	// TODO 切记初始化一个前缀和刚好等于k的。否则回少一个。因为第一次遇到preSum-k时sumCount[0]=0了。从而少了一个从0-i的这个满足情况
	sumCount[0] = 1

	count := 0
	preSum := 0
	for i := 0; i < len(nums); i++ {
		preSum = preSum + nums[i]
		if val, ok := sumCount[preSum-k]; ok {
			count = count + val
		}
		sumCount[preSum]++
	}
	return count
}

// 相交链表。分别跑完两个链表，最终要相等
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB

	for pA != pB {
		if pA.Next != nil {
			pA = pA.Next
		} else {
			pA = headB
		}
		if pB.Next != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}
	return pA
}

// 反转链表 TODO 如果是双向链表呢？
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// 反转双向链表 TODO 理解不是很透彻
func reverseDoubleList(head *DoubleListNode) *DoubleListNode {
	var pre *DoubleListNode
	cur := head

	for cur != nil {
		next := cur.Next

		cur.Next = cur.Prev
		cur.Prev = next
		pre = cur

		cur = next
	}

	return pre
}

// 找终点
func endOfFirstHalf(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	midNode := endOfFirstHalf(head)
	revertSufix := reverseList(midNode.Next)

	p1 := head
	p2 := revertSufix
	result := true
	for result && p2 != nil {
		if p1 != p2 {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return result
}

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// 二叉树的最大深度，递归实现
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDep := maxDepth(root.Left)
	rightDep := maxDepth(root.Right)

	return max(leftDep, rightDep) + 1
}

// 二叉树的最大深度，队列实现，每一次遍历就把所有的叶子节点放进去，遍历时严格控制每次遍历的数量。每遍历完成一次，深度就加一层
func maxDepth2(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	depth := 0

	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			// 弹出节点
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}

// 反转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	root.Left = invertTree(root.Right)
	root.Right = invertTree(root.Left)

	return root

}

// 对称二叉树 TODO 不熟练
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root.Left, root.Right)
}

func check(left *TreeNode, right *TreeNode) bool {
	// 情况 1：都为空，对称
	if left == nil && right == nil {
		return true
	}
	// 情况 2：其中一个为空，另一个不为空，不对称
	if left == nil || right == nil {
		return false
	}
	// 情况 3：值不相等，不对称
	if left.Val != right.Val {
		return false
	}

	// 情况 4：值相等，则继续向下递归检查：
	// 1. 左树的左节点 匹配 右树的右节点
	// 2. 左树的右节点 匹配 右树的左节点
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}

// 对称二叉树 版本2 队列
func isSymmetric2(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)

	for len(queue) > 0 {
		a := queue[0]
		b := queue[1]
		queue = queue[2:]
		if a == nil && b == nil {
			continue
		}
		if a != nil || b != nil || a.Val != b.Val {
			return false
		}
		queue = append(queue, a.Left, b.Right)
		queue = append(queue, a.Right, b.Left)
	}
	return true
}

// 二叉树的直径 TODO 不熟练
func diameterOfBinaryTree(root *TreeNode) int {
	maxZhiJing := 0

	var maxDepth func(*TreeNode) int
	maxDepth = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		leftDepth := maxDepth(tn.Left)
		rightDepth := maxDepth(tn.Right)

		maxZhiJing = max(leftDepth+rightDepth, maxZhiJing)
		return max(leftDepth, rightDepth) + 1
	}
	maxDepth(root)
	return maxZhiJing
}

// 买卖股票的最佳时机
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy := -prices[0]
	sell := 0

	for i := 0; i < len(prices); i++ {
		buy = max(buy, -prices[i])
		sell = max(sell, prices[i]+buy)
	}
	return sell
}

// 买卖股票的最佳时机 交易两次
func maxProfitTwoTransport(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buyFir := -prices[0]
	sellFir := 0
	buySec := -prices[0]
	sellSec := 0

	for i := 0; i < len(prices); i++ {
		p := prices[i]
		buyFir = max(-p, buyFir)
		sellFir = max(sellFir, buyFir+p)
		buySec = max(sellFir-p, buySec)
		sellSec = max(sellSec, prices[i]+buySec)
	}
	return sellSec
}
