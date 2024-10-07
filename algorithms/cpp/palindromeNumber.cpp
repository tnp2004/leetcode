#include<iostream>

using namespace std;

class Solution {
public:
    bool isPalindrome(int x) {
        if (x < 0) return false;
        if (x >= 0 && x < 10) return true;
        if (x % 10 == 0) return false;
        int v = 0;
        int c = x;

        while(c > v) {
            v *= 10;
            v += c % 10;
            c = c / 10;
        }

        if ((c == v) || c == (v / 10)) {
            return true;
        }

        return false;
    }
};