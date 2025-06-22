class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        words = list(filter(None, s.split(" ")))

        return len(words[len(words) - 1])


def main():
    s = "   fly me   to   the moon  "

    sol = Solution()
    print(sol.lengthOfLastWord(s))


if __name__ == "__main__":
    main()

