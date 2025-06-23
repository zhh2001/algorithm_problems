from typing import List


class Solution:
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        word_len = len(words[0])
        words_len = len(words)
        first_letter = {word[0] for word in words}
        index_list = []

        if word_len * words_len > len(s):
            return []

        if word_len * words_len == len(s):
            for word in words:
                s = s.replace(word, "", 1)
            return [] if s else [0]

        for i, ch in enumerate(s):
            words_copy = words.copy()
            if ch in first_letter:
                flag = True
                substr_list = [s[i + word_len * n:i + word_len * (n + 1)] for n in range(words_len)]
                for substr in substr_list:
                    if substr in words_copy:
                        words_copy.remove(substr)
                    else:
                        flag = False
                        break
                if flag:
                    index_list.append(i)
        return index_list


def main():
    solution = Solution()
    print(solution.findSubstring("barfoothefoobarman", ["foo", "bar"]))
    print(solution.findSubstring("wordgoodgoodgoodbestword", ["word", "good", "best", "word"]))
    print(solution.findSubstring("barfoofoobarthefoobarman", ["bar", "foo", "the"]))


if __name__ == '__main__':
    main()
