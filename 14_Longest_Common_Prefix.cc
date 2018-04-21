#include <string>
#include <vector>
#include <iostream>

using namespace std;

class Solution {
    public:
    string longestCommonPrefix(vector<string>& strs) {
        if (strs.empty())
            return "";
        int tag = 0;
        std::string result;
        for (string::size_type i = 0; i < strs[tag].size(); ++i) {
            char c = strs[tag][i];
            for (vector<string>::size_type j = 0; j < strs.size(); ++j) {
                if (strs[j].size() >= i && c == strs[j][i])
                    continue;
                else
                    return result;
            }
            result.push_back(c);
        }
        return result;
    }
};

int main()
{
    std::vector<string> testVect{"flower", "flow", "flight"};
    std::vector<string> testVect1{"dog", "racecar", "car"};
    Solution s;
    cout << s.longestCommonPrefix(testVect);
    cout << s.longestCommonPrefix(testVect1);
    return 0;
}
