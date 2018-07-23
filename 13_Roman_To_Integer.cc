#include <iostream>
#include <string>
#include <map>
#include <vector>
#include <tuple>

using namespace std;

const map<char, int> gMappingTable{{'I', 1}, {'V', 5}, {'X', 10}, {'L', 50}, {'C', 100}, {'D', 500}, {'M', 1000}};

class Solution {
    public:
    int romanToInt(string s) {
        if(s.length() == 1) {
            auto result = gMappingTable.find(s[0]);
            if (result != gMappingTable.end())
                return result->second;
            else
                return 0;
        }
        std::vector<int> intVect;
        for (string::size_type first = 0, second = 1; first < s.length(); ++first, ++second) {
            auto firstInt = gMappingTable.find(s[first])->second;
            auto secondInt = gMappingTable.find(s[second])->second;
            if (firstInt >= secondInt) {
                intVect.push_back(firstInt);
            } else {
                intVect.push_back(secondInt - firstInt);
                ++first;
                ++second;
            }
            
        }
        if (intVect.front() >= intVect.back()) {
            int result = 0;
            for (const auto& val : intVect) {
                result += val;
            }
            return result;
        } else {
            int result = intVect.back();
            for (auto iter = intVect.begin(); iter != intVect.end() - 1; ++iter) {
                result -= *iter;
            }
            return result;
        }
    }
    
};

int main()
{
    Solution s;
    cout << s.romanToInt("III") << endl;
    cout << s.romanToInt("IX") << endl;
    cout << s.romanToInt("IV`") << endl;
    cout << s.romanToInt("MCMXCIV") << endl;
    return 0;
}

