#include <iostream>
#include <unordered_map>
using namespace std;
string minWindow(string s, string t)
{
    unordered_map<char, int> mp;
    if (s.size() < t.size())
        return "";
    if (t.size() == 1 && t[0] == s[0])
        return t;
    for (int i = 0; i < t.size(); i++)
        mp[t[i]]++;
    int start = 0, end = 0;
    int minlen = INT_MAX;
    int left = 0, right = 0;
    bool flag = false;
    unordered_map<char, int> mp2;
    while (end < s.size())
    {
        mp2[s[end]]++;
        flag = true;
        for (auto t : mp)
        {
            if (mp2[t.first] < t.second)
            {
                flag = false;
                break;
            }
        }
        if (flag)
        {
            while (start <= end)
            {
                if (end - start + 1 < minlen)
                {
                    minlen = end - start + 1;
                    left = start;
                    right = end;
                }
                mp2[s[start]]--;
                start++;
                if (mp2[s[start - 1]] < mp[s[start - 1]])
                    break;
            }
        }
        end++;
    }
    if (left + right != 0)
        return s.substr(left, right - left + 1);
    return "";
}
int main()
{
    string s = "ab";
    string t = "a";
    cout << minWindow(s, t);
    return 0;
}