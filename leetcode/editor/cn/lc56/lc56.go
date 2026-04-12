/**
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回
一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。



 示例 1：


输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].


 示例 2：


输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。



 提示：


 1 <= intervals.length <= 10⁴
 intervals[i].length == 2
 0 <= starti <= endi <= 10⁴


 Related Topics 数组 排序 👍 2532 👎 0

*/

package main

import (
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	// 1. 先进行排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	// 2. 合并区间
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		curr := intervals[i]

		if curr[0] <= last[1] { // 有重叠，合并
			if curr[1] > last[1] {
				last[1] = curr[1] // 更新最后一个区间的结束点
			}
		} else { // 无重叠，直接加入
			merged = append(merged, curr)
		}
	}

	return merged
}

// leetcode submit region end(Prohibit modification and deletion)
