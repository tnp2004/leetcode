#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
    public: vector<int> twoSum(vector<int> &nums, int target) {
        unordered_map<int, int> hm;
        vector<int> result;

        for (int i = 0; i < nums.size(); i++) {
            int diff = target - nums[i];

            // check found ?
            if (hm.find(diff) != hm.end())
            {
                // found
                return {i, hm[diff]};
            }
            
            // not found add to hm
            hm[nums[i]] = i;
        }

        return {};
    }
};