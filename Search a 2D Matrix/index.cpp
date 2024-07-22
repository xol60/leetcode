#include <iostream>
#include <vector>
using namespace std;
bool searchMatrix(vector<vector<int>> &matrix, int target)
{
    int n = matrix[0].size(), m = matrix.size();
    int mid, res;
    int left = 0, right = (n * m) - 1;
    while (left <= right)
    {
        mid = left + (-left + right) / 2;
        res = matrix[mid / n][mid % n];
        if (res == target)
            return true;
        else if (res > target)
            right = mid - 1;
        else
            left = mid + 1;
    }
    return false;
}
int main()
{
    vector<vector<int>> matrix = {{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}};
    cout << searchMatrix(matrix, 15);
    return 0;
}