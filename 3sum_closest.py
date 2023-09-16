class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        min_diff = sum_ans = sys.maxsize
        for i in range(len(nums)-1):
            for j in range(i+1, len(nums)):
                left, right = j+1, len(nums)
                while left < right:
                    mid = left + (right - left) // 2
                    temp_sum = nums[i] + nums[j] + nums[mid]
                    if abs(temp_sum - target) < min_diff:
                        min_diff = abs(temp_sum - target)
                        sum_ans = temp_sum
                    if temp_sum == target:
                        return target
                    if temp_sum > target:
                        right = mid
                    else:
                        left = mid + 1
        return sum_ans
