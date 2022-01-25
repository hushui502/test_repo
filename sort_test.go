package test_repo

import "testing"

func TestSort(t *testing.T) {
	nums := []int{1, 2, 3}

	result := []int{1, 2, 3}

	sort_std(nums)

	for i := range nums {
		if nums[i] != result[i] {
			t.Fatal("err!!!!!!!")
		}
	}
}
