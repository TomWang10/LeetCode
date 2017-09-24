#include <vector>
#include <algorithm>
#include <iostream>

using namespace std;

class Solution
{
    public:
        double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2)
        {
            vector<int> v(nums1.size() + nums2.size());
            merge(nums1.begin(), nums1.end(), nums2.begin(), nums2.end(), v.begin());
            if (v.size() % 2 == 0)
            {
                return static_cast<double>((v[v.size()/2 - 1] + v[v.size()/2])) / static_cast<double>(2);
            }
            else
            {
                return v[(v.size() + 1) / 2];
            }
        }
};

int main()
{
    Solution s;
    vector<int> test1 = {1, 2};
    vector<int> test2 = {3, 4};
    cout << s.findMedianSortedArrays(test1, test2) << std::endl;
    return 0;
}

