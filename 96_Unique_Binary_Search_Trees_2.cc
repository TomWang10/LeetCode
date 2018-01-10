/**
 *  * Definition for a binary tree node.
**/
#include <iostream>
#include <vector>

using namespace std;


struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
    public:
            vector<TreeNode*> Helper(int start, int end) {
                vector<TreeNode *> result;
                if (start > end) {
                    result.push_back(NULL);
                    return result;
                }
                for (int i = start; i <= end; ++i) {
                    vector<TreeNode *> left = Helper(start, i - 1);
                    vector<TreeNode *> right = Helper(i + 1, end);
                    for (int j = 0; j < left.size(); ++j) {
                        for (int k = 0; k < right.size(); ++k) {
                            TreeNode *newNode = new TreeNode(i);
                            newNode->left = left[j];
                            newNode->right = right[k];
                            result.push_back(newNode);
                        }
                    }
                }
                return result;
            }

            vector<TreeNode*> generateTrees(int n) {
                if (n == 0) return vector<TreeNode*>();
                return Helper(1, n);            
            }
};

int main()
{
    Solution s;
    std::cout << s.generateTrees(0).size() << std::endl;
    return 0;
}
