package main
func permute(nums []int) [][]int {
    var res [][]int
    var getAllPermutations func(int)

    getAllPermutations = func(idx int) {
        // Base case
        if idx == len(nums) {
            // Create deep copy (Golang sepecific?)
            temp := make([]int, len(nums))
            copy(temp, nums)
            res = append(res, temp)
        }

        for i := idx; i < len(nums); i++ {
            nums[idx], nums[i] = nums[i], nums[idx]
            getAllPermutations(idx+1)
            // Backtracking
            nums[idx], nums[i] = nums[i], nums[idx]
        }
    }

    getAllPermutations(0)
    return res
}
