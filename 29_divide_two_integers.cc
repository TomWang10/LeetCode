#include <string>
#include <iostream>
#include <limits>
#include <cstdlib>

class Solution {
    public:
    int divide(int dividend, int divisor) {
        if (!divisor || (dividend == std::numeric_limits<int>::min() && divisor == -1)) {
            return std::numeric_limits<int>::max();
        }
        int sign = ((dividend < 0) ^ (divisor < 0)) ? -1 : 1;
        long long dvd = labs(dividend);
        long long dvs = labs(divisor);
        int res = 0;
        while (dvd >= dvs) {
            long long temp  = dvs, multiple = 1;
            while (dvd >= (temp << 1)) { 
                temp <<= 1;
                multiple <<= 1;
            }
            dvd -= temp;
            res += multiple;
        }
        return sign == 1 ? res : -res;
    }
};

int main()
{
    Solution s;
    std::cout << std::numeric_limits<int>::min() << std::endl;
    std::cout << s.divide(-2147483648, 1) << std::endl;
    std::cout << s.divide(-2147483648, -1) << std::endl;
    std::cout << s.divide(10, 3) << std::endl;
    std::cout << s.divide(10, -3) << std::endl;
    std::cout << s.divide(10, 10) << std::endl;
}
