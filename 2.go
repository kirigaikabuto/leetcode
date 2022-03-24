package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	total1StrNormal := ""
	total2StrNormal := ""
	for l1 != nil {
		total1StrNormal += strconv.Itoa(l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		total2StrNormal += strconv.Itoa(l2.Val)
		l2 = l2.Next
	}
	total1Str := ""
	total2Str := ""
	for i := len(total1StrNormal) - 1; i >= 0; i-- {
		total1Str += string(total1StrNormal[i])
	}
	for i := len(total2StrNormal) - 1; i >= 0; i-- {
		total2Str += string(total2StrNormal[i])
	}
	number1 := new(big.Int)
	number2 := new(big.Int)
	total1, _ := number1.SetString(total1Str, 10)
	total2, _ := number2.SetString(total2Str, 10)
	total1.Add(total1, total2)
	totalStrNormal := total1.String()
	totalStr := ""
	for i := len(totalStrNormal) - 1; i >= 0; i-- {
		totalStr += string(totalStrNormal[i])
	}
	head := &ListNode{}
	cur := head
	for i, v := range totalStr {
		n, _ := strconv.Atoi(string(v))
		cur.Val = n
		if i != len(totalStr)-1 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return head
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	s := addTwoNumbers(l1, l2)
	current := s
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}
