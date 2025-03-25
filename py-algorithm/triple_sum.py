from typing import List

def triplet_sum(nums: List[int]) -> List[List[int]]:
    output = []
    return output



if __name__ == '__main__':
    nums = [-1, 0, 1, 2, -1, -4]
    print(triplet_sum(nums))  # [[-1, -1, 2], [-1, 0, 1]]
    
    nums = []
    print(triplet_sum(nums))  # []
    
    nums = [0]
    print(triplet_sum(nums))  # []
    
    nums = [0, -1, 2, -3, 1]
    print(triplet_sum(nums))  # [[-3, 1, 2], [-1, 0, 1]]