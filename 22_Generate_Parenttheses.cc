#include <string>
#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    vector<string> generateParenthesis(int n) {
	vector<string> result;
	Helper(result, "", 0, 0, n);
	return result;
    }

    void Helper(vector<string>& result, string curr, int open, int close, int max) {
	if (curr.size() == max * 2) {
	    result.emplace_back(curr);
	    return;
	}
	if (open < max) {
	    Helper(result, curr + "(", open+1, close, max);
	}
	if (close < open) {
	    Helper(result, curr + ")", open, close+1, max);
	}
    }
};

int main()
{
    Solution s;
    auto result = s.generateParenthesis(3);
    for (auto val : result) {
	std::cout << val << std::endl;
    }
    return 0;
}
