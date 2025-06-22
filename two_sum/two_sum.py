from typing import List, Tuple

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i, _ in enumerate(nums):
            for j in range(i+1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]

        return []

nums = [2,7,11,15]
target = 9

s = Solution()
r = s.twoSum(nums, target)

print(r)
