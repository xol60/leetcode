package main
type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res:=&ListNode{};
	cur:=res;
	t:=0;
    for l1!=nil&&l2!=nil{
        cur.Next=&ListNode{};
		cur=cur.Next;
		t=l1.Val+l2.Val+t;
		if t>=10{
			cur.Val=t-10;
			t=1;
		} else {
			cur.Val=t;
			t=0;
		}
		l1=l1.Next;
		l2=l2.Next;
	}
    if l1==nil{
		l1=l2;
	}
	for l1!=nil{
        cur.Next=&ListNode{};
		cur=cur.Next;
		t=l1.Val+t;
		if t>=10{
			cur.Val=t-10;
			t=1;
		} else {
			cur.Val=t;
			t=0;
		}
		l1=l1.Next;
	}
    if t!=0{
        cur.Next=&ListNode{};
		cur=cur.Next;
        cur.Val=t;
    }
  //  fmt.Println(cur.Val)
	return res.Next;
}