package main
//
//import "fmt"
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func removeElements(head *ListNode, val int) *ListNode {
//	cur := head
//	if head == nil {
//		return nil
//	}
//	if head.Next == nil && head.Val == val {
//		return nil
//	} else if head.Next == nil {
//		return head
//	}
//	for cur.Next != nil {
//		if cur.Next.Val == val {
//			cur.Next = cur.Next.Next
//		} else {
//			cur = cur.Next
//		}
//	}
//	if head.Next == nil && head.Val == val {
//		return nil
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
//	response := removeElements(head, 1)
//	fmt.Println(response)
//}
