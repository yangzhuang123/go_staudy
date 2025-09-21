package gostudy

import "sort"

func merge(intervals [][]int) [][]int {
	// 先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var reuslt [][]int

	// 记录第一个区间的开始和结束
	// 若第二个区间的开始小于第一个区间的结束
	// 合并区间，并更新第一个区间的结束
	// 若第二个区间的开始大于第一个区间的结束，则直接添加进结果开始下一个合并

	tempResult := []int{intervals[0][0], intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		if tempResult[1] >= intervals[i][0] && tempResult[1] < intervals[i][1] {
			tempResult[1] = intervals[i][1]
		} else {
			reuslt = append(reuslt, tempResult)
			tempResult = []int{intervals[i][0], intervals[i][1]}
		}
	}
	reuslt = append(reuslt, tempResult)
	return reuslt
}
