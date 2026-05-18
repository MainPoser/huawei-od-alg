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
	fmt.Println("和为 K 的子数组")
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println("搜索插入位置")
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))

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

// 和为 K 的子数组
func subarraySum(nums []int, k int) int {
	sumCount := make(map[int]int)
	// 初始化：前缀和为 0 出现了 1 次（处理从索引 0 开始就满足条件的情况）
	sumCount[0] = 1

	perSum := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		perSum = perSum + nums[i]
		// 寻找符合条件的前缀和 perSum - k
		if val, ok := sumCount[perSum-k]; ok {
			count += val // 累加该前缀和出现的次数，而不是单纯的 count++
		}

		// 将当前前缀和存入或更新到哈希表中
		sumCount[perSum]++
	}
	return count
}

// 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA := headA
	pB := headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

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

// 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 1. 找到前半部分链表的尾节点（如果是奇数个，指向正中间；偶数个，指向前半部分的最后一个）
	firstHalfEnd := endOfFirstHalf(head)
	// 2. 反转后半部分链表
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// 3. 开始比对
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 4. 恢复链表结构（面试加分项，保护原数据）
	firstHalfEnd.Next = reverseList(secondHalfStart)

	return result
}

// 快慢指针找中点
func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	// 确保 fast 和 fast.Next 都有值，防止空指针
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 环形链表
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)              // 左
		res = append(res, node.Val) // 根
		dfs(node.Right)             // 右
	}

	dfs(root)
	return res
}

// 二叉树的中序遍历 版本2
func inorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	curr := root

	// 当 curr 不为空，或者栈内还有节点时，说明还没遍历完
	for curr != nil || len(stack) > 0 {
		// 1. 一路向左，把左子树全部压栈
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// 2. 左边到头了，弹出栈顶节点（当前子树的根节点）
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 3. 访问根节点
		res = append(res, curr.Val)

		// 4. 转向右子树
		curr = curr.Right
	}

	return res
}

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 递归计算左子树和右子树的深度
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return max(leftDepth, rightDepth) + 1
}

// 二叉树的最大深度2
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		// 当前层的节点数量
		size := len(queue)

		// 循环处理当前层的所有节点，并将它们的子节点入队
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:] // 出队

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 遍历完一层，深度加 1
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
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}

// 是否轴对称 DFS
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 检查左子树和右子树是否镜像对称
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

// 是否轴对称 版本2 BFS
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 用切片模拟队列，每次成对地存入两个需要比较的节点
	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		// 每次取出队头的两个节点
		u := queue[0]
		v := queue[1]
		queue = queue[2:]

		// 如果都为空，说明这一对节点对称，继续检查下一对
		if u == nil && v == nil {
			continue
		}
		// 如果只有一个为空，或者值不相等，说明不对称
		if u == nil || v == nil || u.Val != v.Val {
			return false
		}

		// 成对地将下一层需要比较的节点入队：
		queue = append(queue, u.Left)  // u 的左节点
		queue = append(queue, v.Right) // 匹配 v 的右节点

		queue = append(queue, u.Right) // u 的右节点
		queue = append(queue, v.Left)  // 匹配 v 的左节点
	}

	return true
}

// 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	// 定义一个深度优先搜索函数，返回节点的最大深度
	var maxDepth func(*TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左、右子树的深度
		leftDepth := maxDepth(node.Left)
		rightDepth := maxDepth(node.Right)

		// 【核心】：以当前节点为折返点的路径长度为 leftDepth + rightDepth
		// 实时更新全局最大值
		if leftDepth+rightDepth > maxDiameter {
			maxDiameter = leftDepth + rightDepth
		}

		// 返回当前节点作为子树时的最大深度给它的父节点
		if leftDepth > rightDepth {
			return leftDepth + 1
		}
		return rightDepth + 1
	}

	maxDepth(root)
	return maxDiameter
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 搜索插入位置
func searchInsert(nums []int, target int) int {
	return searchInsertHelp(nums, 0, len(nums)-1, target)
}
func searchInsertHelp(nums []int, l, r, target int) int {
	if l > r {
		return l
	}
	mid := l + (r-l)/2
	// 基础出口 2：找到了直接返回索引
	if nums[mid] == target {
		return mid
	}
	// 严格缩减区间：排除掉已经比较过的 mid
	if nums[mid] < target {
		// target 在右边，左边界推进到 mid + 1
		return searchInsertHelp(nums, mid+1, r, target)
	} else {
		// target 在左边，右边界推进到 mid - 1
		return searchInsertHelp(nums, l, mid-1, target)
	}
}

// 有效括号
func isValid(s string) bool {
	stack := make([]byte, 0)
	stack = append(stack, '0')
	mapp := map[byte]byte{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	mapp2 := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	if len(s) <= 1 {
		return false
	}
	for _, c := range []byte(s) {
		if _, ok := mapp[c]; ok {
			stack = append(stack, c)
		} else {
			pop := stack[len(stack)-1:]
			stack = stack[:len(stack)-1]
			if pop[0] == '0' {
				return false
			}
			if pop[0] != mapp2[c] {
				return false
			}
		}
	}
	if len(stack) > 1 {
		return false
	}
	return true
}

// 爬楼梯
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 爬楼梯 版本2
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}

	// p 对应 n-2，q 对应 n-1，r 对应当前 n
	p, q, r := 1, 2, 0

	for i := 3; i <= n; i++ {
		r = p + q // 当前阶 = 前一阶 + 前两阶
		p = q     // 状态向后移动
		q = r
	}

	return r
}

// 杨辉三角
func generate(numRows int) [][]int {
	c := make([][]int, numRows)
	for i := range c {
		c[i] = make([]int, i+1)
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			// 左上方的数 + 正上方的数
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
	return c
}
