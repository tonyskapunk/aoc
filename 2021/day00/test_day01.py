import unittest

import day01

class TestDay01(unittest.TestCase):
    def test_empty(self):
        measurements = []
        
        count = day01.count_increments(measurements)
        self.assertEqual(count, 0)

    def test_no_increment_single(self):
        measurements = [
            1
        ]
        
        count = day01.count_increments(measurements)
        self.assertEqual(count, 0)

    def test_no_increment_multiple(self):
        measurements = [
            1,
            1,
            1
        ]
        
        count = day01.count_increments(measurements)
        self.assertEqual(count, 0)

    def test_single_increment_single(self):
        measurements = [
            1,
            2
        ]
        
        count = day01.count_increments(measurements)
        self.assertEqual(count, 1)

    def test_single_increment_multiple(self):
        measurements = [
            1,
            2,
            1,
            2
        ]
        
        count = day01.count_increments(measurements)
        self.assertEqual(count, 2)

    def test_full(self):
        measurements = [
          199,
          200,
          208,
          210,
          200,
          207,
          240,
          269,
          260,
          263
        ]
        
        count = day01.count_increments(measurements)
        self.assertEqual(count, 7)

if __name__ == "__main__":
    unittest.main()
