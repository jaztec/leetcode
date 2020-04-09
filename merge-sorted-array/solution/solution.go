package solution

import "sort"

type result []int

func (r result) Len() int           { return len(r) }
func (r result) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r result) Less(i, j int) bool { return r[i] < r[j] }

func merge(nums1 []int, m int, nums2 []int, n int) {
	var r result
	r = append(nums1[:m], nums2...)
	sort.Sort(r)
	nums1 = r
}
