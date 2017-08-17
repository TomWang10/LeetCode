#include <vector>
#include <algorithm>
#include <iostream>

using namespace std;

class Solution
{
    public:
        vector<int> TwoSum(vector<int>& nums, int target)
        {
            for (vector<int>::iterator iter = nums.begin(); iter != nums.end(); ++iter)
            {
                for (vector<int>::iterator iter2 = iter + 1; iter2 != nums.end(); ++iter2)
                {
                    if (*iter + *iter2 == target)
                    {
                        return {static_cast<int>(iter - nums.begin()), 
                            static_cast<int>(iter2 - nums.begin())};
                    }
                }
            }
        }
};

int main()
{
    std::vector<int> nums = {2, 7, 11, 15};
    int target = 9;
    Solution s;
    auto result = s.TwoSum(nums, target);
    std::cout << result[0] << result[1] << std::endl;
    return 0;
}
