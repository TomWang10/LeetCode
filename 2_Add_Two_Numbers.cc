#include <iostream>

struct ListNode {
int val;
ListNode *next;
ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
    public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
            int carryDigit = 0;
            ListNode currNode(-1);
            ListNode* pre = &currNode;
            while (l1 || l2)
            {
               auto newVal = (l1 ? l1->val : 0) + (l2 ? l2->val : 0) + carryDigit; 
               carryDigit = newVal / 10;
               newVal = newVal % 10;
               pre->next = new ListNode(newVal);
               pre = pre->next;
               l1 = l1 ? l1->next : NULL;
               l2 = l2 ? l2->next : NULL;
            }
            if (carryDigit != 0)
            {
                pre->next = new ListNode(carryDigit);
            }
            return currNode.next;
        }
};



int main()
{
    ListNode* l1 = new ListNode(1);    
    l1->next = new ListNode(8);

    ListNode* l2 = new ListNode(0);    
    
    Solution s;
    auto result = s.addTwoNumbers(l1, l2);
    while (result)
    {
        std::cout <<result->val << "\n";
        result = result->next;
    }
    return 0;
}
