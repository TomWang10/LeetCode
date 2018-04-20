#include <iostream>

/**
 *  * Definition for a binary tree node.
 *   * struct TreeNode {
 *    *     int val;
 *     *     TreeNode *left;
 *      *     TreeNode *right;
 *       *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 *        * };
 *         */
 struct TreeNode {
 int val;
 TreeNode *left;
 TreeNode *right;
 TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };

class Solution {
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
	if (q == nullptr && p == nullptr) {
	    return true;
	}
	if (q != nullptr && p != nullptr && p->val == q->val) {
	    return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
	}
	return false;
    }
};

int main()
{
    return 0;
}
