#include <iostream>
#include <string>

using namespace std;


class Solution {
public:
    int strStr(string haystack, string needle) {
        if (needle.empty())
            return -1;
        if (needle.size() > haystack.size()) {
            return -1;
        }
        for (int i = 0; i < haystack.length() - needle.length() + 1; ++i) {
            if (haystack[i] == needle[0]) {
                bool isMatch = true;
                for (int j = 1; j < needle.length(); ++j) {
                    if (haystack[i + j] != needle[j]) {
                        isMatch = false;
                        break;
                    }
                }
                if (isMatch)
                    return i;
            }
        } 
        return -1;
    }
};

int main()
{
    Solution s;
    std::cout << s.strStr("hello", "ll") << std::endl;
    std::cout << s.strStr("hello", "lla") << std::endl;
    return 0;
}
