"""
Given an array of strings, group anagrams together.
"""

from typing import List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagram_dict = {}
        for i, s in enumerate(strs):
            sorted_s = "".join(sorted(s))
            if sorted_s not in anagram_dict:
                anagram_dict[sorted_s] = []
            anagram_dict[sorted_s].append(s)
        return list(anagram_dict.values())
        
    
if __name__ == "__main__":
    sol = Solution()
    strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
    print(sol.groupAnagrams(strs))