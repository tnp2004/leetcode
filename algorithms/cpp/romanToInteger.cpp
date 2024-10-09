#include<iostream>
#include<unordered_map>
#include<string>

using namespace std;

class Solution {
public:
    int romanToInt(string s) {
        unordered_map<char, int> hm = {
            {'I', 1},
            {'V', 5},
            {'X', 10},
            {'L', 50},
            {'C', 100},
            {'D', 500},
            {'M', 1000},
        };
        int sum = 0;
        for (int i = 0; i < s.size(); i++) {
            if (hm[s[i]] < hm[s[i+1]]) {
                sum -= hm[s[i]];
                continue;
            }

            sum += hm[s[i]];
        }

        return sum;
    }
};