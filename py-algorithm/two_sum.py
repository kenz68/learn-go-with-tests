from typing import List

def pair_sum_sorted_v1(nums: List[int], target: int) -> List[int]:
    d = {}
    for i, num in enumerate(nums):
        complement = target - num
        if complement in d:
            return [d.get(complement), i]
        d[num] = i
    return []

def pair_sum_sorted(nums: List[int], target: int) -> List[int]:
    left, right = 0, len(nums) - 1
    while left < right:
        sum = nums[left] + nums[right]
        if sum == target:
            return [left, right]
        elif sum < target:
            left += 1 # move left pointer to the right to increase the sum
        else:
            right -= 1 # move right pointer to the left to decrease the sum
            
    return []
    
if __name__ == '__main__':
    nums = [-5,-2,3,4,6]
    target = 7
    print(pair_sum_sorted(nums, target))  # [0, 1]

# py py-algorithm/two_sum.py 

