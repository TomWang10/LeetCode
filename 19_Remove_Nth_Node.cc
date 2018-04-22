#include <iostream>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
    public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        int ListSize = 0;
        auto tmpHead = head;
        while (tmpHead) {
            ListSize++;
            tmpHead = tmpHead->next;
        }
        decltype(head) preNode = nullptr;
        tmpHead = head;
        auto DeletePos = 0;
        while (ListSize - DeletePos++ != n) {
            preNode = tmpHead;
            tmpHead = tmpHead->next;
        }
        if (preNode == nullptr) {
            head = head->next;
            delete tmpHead;
            return head;
        }
        preNode->next = tmpHead->next;
        delete tmpHead;
        return head;
    }
};

int main()
{
    ListNode* head = new ListNode(1);
    ListNode* curr = head;
    for (int i = 2; i <= 5; ++i) {
        curr->next = new ListNode(i);
        curr = curr->next;
    }
    Solution s;
    head = s.removeNthFromEnd(head, 5);
    while (head) {
        std::cout << head->val << "->";
        head = head->next;
    }
    std::cout << "fuck you";
    return 0;
}
