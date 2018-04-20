#include <string>
#include <iostream>
#include <exception>
#include <limits>
#include <map>

using namespace std;

// 1. +10000, -10000, >INT_MAX , < INT_MIN, empty, 0000100, "   1000", 

class Solution {
    public:
    int myAtoi(string str) {
        int pos = 0,  sign = 1;
        long long base = 0;
        while (str[pos] == ' ')
            ++pos;
        if (str[pos] == '-') {
            sign = -1;
            ++pos;
        } else if (str[pos] == '+') {
            ++pos;
        }
        for (; str[pos] >= '0' && str[pos] <= '9'; ++pos) {
            base = base * 10 + (str[pos] - '0');
            if (base > std::numeric_limits<int>::max()) {
                return sign == 1 ? std::numeric_limits<int>::max() :
                    std::numeric_limits<int>::min();
            }
        }
        return static_cast<int>(base * sign);
    }
};

int main()
{
    Solution s;
    cout << s.myAtoi("2147483648") << endl;
    cout << s.myAtoi("1") << endl;
    cout << s.myAtoi(" +1") << endl;
    cout << s.myAtoi(" -1") << endl;
    cout << s.myAtoi("0000111111") << endl;
    cout << s.myAtoi("+0000111111") << endl;
    return 0;
}
