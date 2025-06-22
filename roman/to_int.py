class Solution:
    def __init__(self) -> None:
        self.nums = {
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M": 1000,

            "IV": 3,
            "IX": 8,
            "XL": 30,
            "XC": 80,
            "CD": 300,
            "CM": 800,
        }

    def romanToInt(self, s: str) -> int:
        sum = 0

        for i, n in enumerate(s):
            key = n

            if i != 0 and self.nums.get(s[i-1]+s[i]):
                key = s[i-1]+s[i]

            found = self.nums[key]

            sum += found

        return sum

def main():
    s = Solution()
    print(s.romanToInt("III"))
    print(s.romanToInt("IV"))
    print(s.romanToInt("LVIII"))
    print(s.romanToInt("MCMXCIV"))

if __name__ == "__main__":
    main()