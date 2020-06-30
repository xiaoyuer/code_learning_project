func removeDuplicates(nums []int) int {
    left, right, size := 1, 1, len(nums)
    if size < 2 {
        return size
    }
    for right < size {
        if nums[right] != nums[right-1] {
            nums[left] = nums[right]
            left++
        }
        right++
    }
    return left
}