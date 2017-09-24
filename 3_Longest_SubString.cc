#include <string>
#include <iostream>
#include <set>

using namespace std;

class Solution
{
    public:
        int lengthOfLongestSubstring(string s)
        {
            int longestLength = 0;
            for (string::size_type index = 0; index < s.size(); ++index)
            {
                if (s.size() - index <= static_cast<unsigned long>(longestLength))
                    return longestLength;
                std::set<char> tmpSet;
                for (auto subIndex = index; subIndex < s.size(); ++subIndex)
                {
                    if (tmpSet.find(s[subIndex]) == tmpSet.end())
                    {
                        tmpSet.insert(s[subIndex]);
                    }
                    else
                    {
                        break;
                    }
                }
                if (tmpSet.size() > static_cast<unsigned long>(longestLength))
                    longestLength = tmpSet.size();
            }
            return longestLength;
        }
};

int main()
{
    Solution s;
    std::string testStr = "aaaaa";
    std::cout << s.lengthOfLongestSubstring(testStr) << std::endl;
    return 0;
}
