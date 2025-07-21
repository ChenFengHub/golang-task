package main

import (
	"errors"
	"fmt"
	"slices"
)

func main() {
	// 控制流程
	// 题目1： 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
	numArr := []int{4, 4, 1, 2, 1, 2, 100}
	matchNum, err := findRepeatOnceNum(numArr)
	if err == nil {
		fmt.Println("题目1:%v中只出现一次的数字是:", numArr, matchNum)
	} else {
		fmt.Printf("题目1:%v数组中未找到只出现一次元素,err:%v\n", numArr, err)
	}
	// 题目2：判断一个整数是否是回文数（回文数是指正序（从左到右）和倒序（从右到左）读都是一样的整数，例如 121、1331、12321）
	// 考察：数字操作、条件判断
	palindTestNum := 12321
	if isPalindrome(palindTestNum) {
		fmt.Printf("题目2:%d 是回文数\n", palindTestNum)
	} else {
		fmt.Printf("题目2:%d 不是回文数\n", palindTestNum)
	}

	// 字符串
	// 题目3：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
	// 考察：字符串处理、栈的使用
	// 思路：遍历字符串，遇到左括号入栈，遇到右括号判断栈顶是否为对应的左括号，是则出栈，否则无效。最后栈为空则有效。
	validTestStr := "([{}])())"
	if isValidStr(validTestStr) {
		fmt.Printf("题目3:字符串 \"%s\" 是有效的括号字符串\n", validTestStr)
	} else {
		fmt.Printf("题目3:字符串 \"%s\" 不是有效的括号字符串\n", validTestStr)
	}
	// 题目4：最长公共前缀-查找字符串数组中的最长公共前缀
	// 考察：字符串处理、循环嵌套
	// 思路：取第一个字符串作为初始前缀；遍历数组中剩下的每个字符串，不断缩短前缀，直到它是当前字符串的前缀为止；如果前缀为空，则没有公共前缀返回""；遍历结束后，剩下的前缀就是最长公共前缀
	commonPrefixArr := []string{"flower", "flow", "flight"}
	commonPrefix := getCommonPrefix(commonPrefixArr)
	fmt.Printf("题目4:%v最长公共前缀是:%v\n", commonPrefixArr, commonPrefix)

	// 基本数据类型
	// 题目5：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
	// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
	// 这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
	// 将大整数加 1，并返回结果的数字数组。
	// 考察：数组操作、进位处理
	bigIntArr := []int{9, 9, 9, 9}
	bigIntArrRes := generateBitIntPlusOne(bigIntArr)
	fmt.Printf("题目5:%v数组加后的数组为:%v\n", bigIntArr, bigIntArrRes)

	// 引用类型：切片
	// 题目6： 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，
	// 使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地
	// 修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i
	// 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j]
	// 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
	nums := []int{1, 1, 2, 3, 3, 4, 5, 5}
	fmt.Print("题目6:原数组:", nums)
	newLength := deduplicateArr(nums)
	fmt.Printf("去重后新数组长度为:%d,新数组为:%v\n", newLength, nums[:newLength])
	// 题目7：合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
	// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，
	// 然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；
	// 如果没有重叠，则将当前区间添加到切片中。
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {17, 20}}
	fmt.Print("题目7:合并前的区间为:", intervals)
	intervalsRes := mergedIntervals(intervals)
	fmt.Printf(" 合并后的区间为:%v\n", intervalsRes)

	// 基础
	// 两数之和
	// 题目8：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
	// 考察：数组遍历、map使用
	numsArr := []int{2, 7, 11, 15}
	target := 9
	fmt.Print("题目8:原数组:", numsArr, "目标值:", target)
	t1, t2, err := getSumTargetTwoElement(numsArr, target)
	if err != nil {
		fmt.Println(" 结果：", err)
	} else {
		fmt.Printf(" 结果：%d + %d = %d\n", t1, t2, target)
	}

}
func getSumTargetTwoElement(nums []int, target int) (int, int, error) {
	// 找出数据中币target小于等于的数
	prepareNums := []int{}
	for _, v := range nums {
		if v <= target {
			prepareNums = append(prepareNums, v)
		}
	}
	if len(prepareNums) <= 0 || (len(prepareNums) == 1 && prepareNums[0] != target) {
		return 0, 0, errors.New("不存在这两个数")
	}

	for i := 0; i < len(prepareNums); i++ {
		for j := i + 1; j < len(prepareNums); j++ {
			if (prepareNums[i] + prepareNums[j]) == target {
				return prepareNums[i], prepareNums[j], nil
			}
		}
	}

	return 0, 0, errors.New("不存在这两个数")
}

func mergedIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 先根据区间第一个大小进行升序排序
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				// 交换位置
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}

	for i := 0; i < len(intervals)-1; i++ {
		// 判断区间重合的，就将两个合并成一个并将后续数据向前移位1
		j := i + 1
		if intervals[i][1] >= intervals[j][0] {
			// 合并区间
			intervals[i][1] = max(intervals[i][1], intervals[j][1])
			// 将后续区间向前移位
			for ; j < len(intervals)-1; j++ {
				intervals[j] = intervals[j+1]
			}
			// 截断数组长度
			intervals = intervals[:len(intervals)-1]
			i-- // 重新检查当前区间与更后一个区间
		}
	}

	return intervals
}

func deduplicateArr(nums []int) int {
	for i := 0; i < len(nums); i++ {
		j := i + 1
		if (nums[i] == nums[j]) && (j < len(nums)) {
			// 重复触发移位操作并截断
			beginIndex := j
			preIndex := i
			for beginIndex < len(nums) {
				nums[preIndex] = nums[beginIndex]
				preIndex++
				beginIndex++
			}
			nums = nums[:len(nums)-1] // 截断数组
		}

	}

	return len(nums)
}

func generateBitIntPlusOne(bigIntArr []int) []int {
	if len(bigIntArr) == 0 {
		return bigIntArr
	}

	resArr := make([]int, len(bigIntArr))

	isPlusOne := true
	for i := len(bigIntArr) - 1; i >= 0; i-- {
		curInt := bigIntArr[i]
		if isPlusOne {
			curInt++
		}
		if curInt == 10 {
			resArr[i] = 0
		} else {
			resArr[i] = curInt
			isPlusOne = false
		}
	}

	if isPlusOne {
		// 仍然有进位则增加最前一位1
		return append([]int{1}, resArr...)
	} else {
		return resArr
	}
}

func getCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		minLen := 0
		if len(prefix) >= len(strs[i]) {
			minLen = len(strs[i])
		} else {
			minLen = len(prefix)
		}

		for minLen > 0 {
			if prefix[:minLen] == strs[i][:minLen] {
				prefix = prefix[:minLen]
				break
			} else {
				minLen--
				if minLen <= 0 {
					prefix = ""
					break
				}
			}
		}

		if minLen <= 0 {
			prefix = ""
		}
	}

	return prefix // 如果前缀为空，直接返回
}

func isValidStr(str string) bool {
	stack := []string{}
	leftChars := []string{"(", "{", "["}
	rightChars := []string{")", "}", "]"}
	matchMap := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	for _, char := range str {
		if slices.Contains(leftChars, string(char)) {
			stack = append(stack, string(char))
		} else if slices.Contains(rightChars, string(char)) {
			if len(stack) > 0 && stack[len(stack)-1] == matchMap[string(char)] {
				stack = stack[:len(stack)-1] // 出栈
			} else {
				stack = append(stack, string(char)) // 入栈
			}
		}
	}
	if len(stack) == 0 {
		return true // 栈为空，说明所有括号都匹配
	} else {
		return false // 栈不为空，说明有未匹配的括号
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	if reversed == original {
		return true
	} else {
		return false
	}
}

func findRepeatOnceNum(nums []int) (int, error) {
	numRepeatMap := make(map[int]int, len(nums))

	for _, v := range nums {
		repertNum, ok := numRepeatMap[v]
		if !ok {
			repertNum = 0
		}
		numRepeatMap[v] = repertNum + 1
	}

	for k, v := range numRepeatMap {
		if v == 1 {
			return k, nil
		}
	}

	return 0, errors.New("没有找到只出现一次的数字")
}
