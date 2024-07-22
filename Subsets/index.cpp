#include<iostream>
#include<vector>
using namespace std;
void backtrack(vector<int>& subset,vector<int>& nums, vector<vector<int>>& result, int start){
    result.push_back(subset);
    for(int i = start; i < nums.size(); i++){
        subset.push_back(nums[i]);
        backtrack(subset,nums,result, i+1);
        subset.pop_back();
    }
}
vector<vector<int>> subsets(vector<int>& nums) {
        vector<vector<int>> result;
        vector<int> subset;
        backtrack(subset,nums,result, 0);
        return result;
}
int main(){
    return 0;
}