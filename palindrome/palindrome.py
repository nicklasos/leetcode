class Solution:
    def isPalindrome(self, x: int) -> bool:
        s = str(x)
        ss = s[::-1]

        return s == ss


def main():
    x = 121
    
    s = Solution()
    r = s.isPalindrome(x)

    print(r)

if __name__ == "__main__":
    main()