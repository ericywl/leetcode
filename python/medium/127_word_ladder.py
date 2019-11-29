"""
Given two words (beginWord and endWord), and a dictionary's word list, 
find the length of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
Note:

Return 0 if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
"""


import collections
from typing import List


class Solution:
    def ladderLength(self, begin_word: str, end_word: str, word_list: List[str]) -> int:
        def create_generic(word: str, i: int) -> str:
            return word[:i] + "_" + word[i + 1:]
        
        def build_dict(word_list: List[str]) -> dict:
            res = collections.defaultdict(list)
            for word in word_list:
                for i in range(len(word)):
                    generic = create_generic(word, i)
                    res[generic].append(word)
            return res
        
        if end_word not in word_list:
            return 0

        generic_dict, visited = build_dict(word_list), set()
        queue = collections.deque([(begin_word, 1)])
        while queue:
            curr_word, num_transform = queue.popleft()
            if curr_word in visited:
                continue
            if curr_word == end_word:
                return num_transform
            visited.add(curr_word)
            for i in range(len(curr_word)):
                generic = create_generic(curr_word, i)
                for next_word in generic_dict[generic]:
                    if next_word not in visited:
                        queue.append((next_word, num_transform + 1))
        return 0


if __name__ == "__main__":
    sol = Solution()
    begin_word = "hit"
    end_word = "cog"
    word_list = ["dot", "hot", "dog", "lot", "log", "cog"]
    print(sol.ladderLength(begin_word, end_word, word_list))
