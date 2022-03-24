package main
//
//import "fmt"
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func deleteDuplicates(head *ListNode) *ListNode {
//	cur := head
//	if head == nil {
//		return nil
//	}
//	if head.Next == nil {
//		return head
//	}
//	for cur.Next != nil {
//		if cur.Next.Val == cur.Val {
//			cur.Next = cur.Next.Next
//		} else {
//			cur = cur.Next
//		}
//	}
//	return head
//}
//
//func main() {
//	second := &ListNode{
//		Val:  1,
//		Next: nil,
//	}
//	first := &ListNode{
//		Val:  1,
//		Next: second,
//	}
//	head := &ListNode{}
//	head.Val = 3
//	head.Next = first
//	response := deleteDuplicates(head)
//	for response != nil {
//		fmt.Println(response.Val)
//		response = response.Next
//	}
//}
