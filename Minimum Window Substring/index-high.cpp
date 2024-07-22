#include<iostream>
#include<unordered_map>
using namespace std;
string minWindow(string s, string t)
{
    unordered_map<char, int> mp;
    for(auto temp:t)
        mp[temp]++;
    int start = 0;
    int min=INT_MAX;
    int count=0;
    int left = 0, right = 0;
    for(int end=0;end<s.size();end++){
        if(mp.find(s[end])!=mp.end()){
            mp[s[end]]--;
            if(mp[s[end]]>=0)
            count++;
        }
        while(count==t.size()){
            if(min>end-start+1){
                min = end-start+1;
                left = start;
                right = end;
            }
            if(mp.find(s[start])!=mp.end()){
                mp[s[start]]++;
                if(mp[s[start]]>0)
                    count--;
            }
            start++;
        }
    }
    
    return min==INT_MAX?"":s.substr(left, right-left+1);
}
int main()
{
    string s = "ADOBECODEBANC";
    string t = "ABC";
    cout << minWindow(s, t);
    return 0;
}