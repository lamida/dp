import unittest
from lecture3 import nSum

class TestNSum(unittest.TestCase):

    def test_nSum(self):
        expected = [0, 1, 3, 6, 10, 15, 21, 28, 36, 45]
        for i in range(len(expected)):
            self.assertEqual(expected[i], nSum(i))

if __name__ == "__main__":
    unittest.main()