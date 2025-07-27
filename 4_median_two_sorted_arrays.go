package main
import "math"
func maxFloat(a, b int) float64 {
    if a > b {
        return float64(a)
    }
    return float64(b)
}
func minFloat(a,b int )float64{
    if a<b{
        return float64(a)
    }
    return float64(b)
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n1,n2:=len(nums1),len(nums2);
    if n1>n2{
        return findMedianSortedArrays(nums2, nums1)
    }
    low,high:=0,n1;
    l1,l2,r1,r2:=0,0,0,0;
    m1,m2:=0,0;
    left:=(n1+n2+1)/2;
    for low<=high{
        m1=(low+high)/2;
        m2=left-m1;
        l1,l2,r1,r2=math.MinInt64, math.MinInt64, math.MaxInt64, math.MaxInt64
        if m1<n1{
            r1=nums1[m1]
        }  
        if m2<n2{
            r2=nums2[m2]
        } 
        if m1>=1{
            l1=nums1[m1-1]
        } 
        if m2>=1{
            l2=nums2[m2-1]
        }
        if l1<=r2&&l2<=r1{
            if (n1+n2)%2==1{
                //fmt.Println("result",l1,l2)
                return maxFloat(l1,l2)
            } else{
                return (maxFloat(l1,l2)+minFloat(r1,r2))/2
            }
        } else if l1>r2 {
            high = m1-1
        } else {
            low = m1+1
        }
    }
    return 0;
}