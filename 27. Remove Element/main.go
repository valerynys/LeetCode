package main

func removeElement(nums []int, val int) int {
	head := 0
	tail := len(nums) - 1

	for head <= tail {
		for head <= tail && nums[head] != val {
			head++
		}
		if head > tail {
			nums = nums[:head]
			break
		}
		for head <= tail && nums[tail] == val {
			tail--
		}
		if head > tail {
			nums = nums[:head]
			break
		}
		nums[head], nums[tail] = nums[tail], nums[head]
	}

	return len(nums)
}

func main() {

}
