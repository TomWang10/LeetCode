#include <string>
#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution
{
public:
    string longestPalindrome(string s)
    {
        int start = 0, length = 0;
        
        for (int index = 0; index <= s.length(); ++index)
        {
            if (s.length() - index < length)
                break;
            for (int rightIndex = s.length() - 1; rightIndex >= index; --rightIndex)
            {
                if (s[index] == s[rightIndex])
                {
                    auto tmpIndex = index;
                    auto tmpRightIndex = rightIndex;
                    ++tmpIndex;
                    --tmpRightIndex;
                    for (; s[tmpIndex] == s[tmpRightIndex] && tmpIndex < tmpRightIndex; ++tmpIndex, --tmpRightIndex);
                    // check whether found
                    if (tmpIndex >= tmpRightIndex)
                    {
                        if (length < (rightIndex - index + 1))
                        {
                            length = rightIndex - index + 1;
                            start = index;
                            break;
                        }
                    }
                }

            }
        }
        if (length == 0)
            return "";
        else
        {
            return string(s, start, length);
        }
    }
};

class Solution2
{
public:
    string longestPalindrome(string s)
    {
        std::string tmpStr= "$#";
        for (auto index = 0; index < s.length(); ++index)
        {
            tmpStr += s[index];
            tmpStr += "#";
        }
        vector<int> maxTable(tmpStr.size(), 0);
        int currMaxIndex = 0, currMaxRight = 0, maxRadius = 0;
        for (auto i = 0; i < tmpStr.length(); ++i)
        {
           maxTable[i] = currMaxRight > i ?  min(maxTable[currMaxIndex * 2 - i], currMaxRight - i) : 1;
           while(tmpStr[i + maxTable[i]] == tmpStr[i - maxTable[i]]) ++maxTable[i];
           if ((currMaxRight - currMaxIndex) < maxTable[i])
           {
               currMaxIndex = i;
               currMaxRight = i + maxTable[i];
               maxRadius = maxTable[i];
           }
        }
        
        return string(s, (currMaxIndex - maxRadius) / 2, maxRadius - 1);
    }
};


int main()
{
    Solution2 s;
    string testStr = "cbbd";
    string testStr1 = "$#c#b#b#d";
    cout << s.longestPalindrome(testStr) << "\n";
    return 0;
}

