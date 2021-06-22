package leetcode

func rotate(nums []int, k int) {
	size_ := len(nums)
	k = k % size_
	cnt := 0
	for start := 0; cnt < size_; start++ {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % size_
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			cnt++
			if start == current {
				break
			}
		}
	}
}
