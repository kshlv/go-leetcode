// Package mean implements _an optimal_ [O(log(m+n))] search for
// two sorted arrays' median
//
// Original problem:
// https://leetcode.com/problems/median-of-two-sorted-arrays/
package mean

func findMedianSortedArray(nums []int) float64 {
	l := len(nums)
	if l == 0 {
		return 0.0
	}

	if l%2 == 0 {
		return float64(nums[l/2-1]+nums[l/2]) / 2.0
	}

	return float64(nums[(l-1)/2])
}

// Mean has an actual signature of findMedianSortedArrays(nums1 []int, nums2 []int) float64
// I'm keeping it simple and more ideomatic
func Mean(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	if m == 0 {
		return findMedianSortedArray(nums2)
	}
	if n == 0 {
		return findMedianSortedArray(nums1)
	}

	if nums1[0] > nums2[0] {
		nums1, nums2 = nums2, nums1
		m, n = len(nums1), len(nums2)
	}

	result := []int{nums1[0]}
	i, j := 1, 0

	for {
		if i == m {
			result = append(result, nums2[j:]...)
		}
		if j == n {
			result = append(result, nums1[i:]...)
		}

		if len(result) > (m+n)/2 {
			if (m+n)%2 == 0 {
				return float64(result[(m+n)/2-1]+result[(m+n)/2]) / 2.0
			}
			return float64(result[(m+n-1)/2])
		}

		if i >= m ||
			(nums2[j] >= result[len(result)-1] && nums2[j] <= nums1[i]) {
			result = append(result, nums2[j])
			if j < n {
				j++
			}
		}

		if j >= n ||
			(nums1[i] >= result[len(result)-1] && nums1[i] <= nums2[j]) {
			result = append(result, nums1[i])
			if i < m {
				i++
			}
		}
	}
}
