#include <iostream>

/**
 *  * Definition for singly-linked list.
*/
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

void PrintList(ListNode* head) {
    while (head) {
        std::cout << head->val << "->";
        head = head->next;
    }
    std::cout << std::endl;
}

class Solution {
    public:
    ListNode* swapPairs(ListNode* head) {
            if (head == nullptr || head->next == nullptr)
                return head;
            ListNode* pre = nullptr;
            ListNode* first = head;
            ListNode* second = head->next;
            ListNode* pos = second->next;
            ListNode* result = head->next;
            while (true) {
                if (pre != nullptr)
                    pre->next = second;
                second->next = first;
                first->next = pos;
                if (pos == nullptr || pos->next == nullptr) {
                    return result;
                }
                pre = first;
                first = pos;
                second = first->next;
                pos = second->next;
                PrintList(result);
            }
            return result;
        }
};


int main()
{
    ListNode* head = new ListNode(1);
    ListNode* tmpHead = head;
    for (int i = 2; i <= 3; ++i) {
       tmpHead->next = new ListNode(i); 
       tmpHead = tmpHead->next;
    }
    Solution s;
    PrintList(head);
    ListNode* result = s.swapPairs(head);
    PrintList(result);
    return 0;
}
