import unittest
from two_sum import pair_sum_sorted

class TestPairSumSorted(unittest.TestCase):
    
    def test_pair_sum_sorted_found(self):
        self.assertEqual(pair_sum_sorted([-5, -2, 3, 4, 6], 7), [2, 3])
        self.assertEqual(pair_sum_sorted([1, 2, 3, 4, 5], 9), [3, 4])
        self.assertEqual(pair_sum_sorted([0, 1, 2, 3, 4], 3), [1, 2])
    
    def test_pair_sum_sorted_not_found(self):
        self.assertEqual(pair_sum_sorted([1, 2, 3, 4, 5], 10), [])
        self.assertEqual(pair_sum_sorted([0, 1, 2, 3, 4], 8), [])
    
    def test_pair_sum_sorted_empty_list(self):
        self.assertEqual(pair_sum_sorted([], 5), [])
    
    def test_pair_sum_sorted_single_element(self):
        self.assertEqual(pair_sum_sorted([5], 5), [])
    
    def test_pair_sum_sorted_duplicates(self):
        self.assertEqual(pair_sum_sorted([1, 2, 3, 3, 4], 6), [2, 3])
        self.assertEqual(pair_sum_sorted([1, 1, 1, 1, 1], 2), [0, 1])

if __name__ == '__main__':
    unittest.main()