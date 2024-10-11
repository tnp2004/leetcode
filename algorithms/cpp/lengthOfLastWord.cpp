#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
    int lengthOfLastWord(string s) {
        int last = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s[i] != ' ') {
                int len = 0;
                for (int j = i; j < s.length(); j++) {
                    if (s[j] == ' ') break;
                    len++;
                    last = len;
                }
                i += len;
            }
        }

        return last;
    }
};

int main() {
    Solution s;
    cout << s.lengthOfLastWord("   fly me   to   the moon  ");

    return 0;
}