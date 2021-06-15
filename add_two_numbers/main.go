package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node3A := ListNode{
		Val:  9,
		Next: nil,
	}

	node2A := ListNode{
		Val:  9,
		Next: &node3A,
	}

	node1A := ListNode{
		Val:  9,
		Next: &node2A,
	}

	node3B := ListNode{
		Val:  9,
		Next: nil,
	}

	node2B := ListNode{
		Val:  9,
		Next: &node3B,
	}

	node1B := ListNode{
		Val:  9,
		Next: &node2B,
	}

	node1C := addTwoNumbers(&node1A, &node1B)
	print(node1C)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	var l3 *ListNode = nil
	var output *ListNode = nil
	value := 0

	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			value = l1.Val + l2.Val + carry
			carry = value / 10
			value = value % 10

			if l3 == nil {
				l3 = &ListNode{
					Val:  value,
					Next: nil,
				}
				output = l3
			} else {
				l3.Next = &ListNode{
					Val:  value,
					Next: nil,
				}
				l3 = l3.Next
			}
			l1 = l1.Next
			l2 = l2.Next

		} else if l1 != nil {
			value = l1.Val + carry
			carry = value / 10
			value = value % 10

			l3.Next = &ListNode{
				Val:  value,
				Next: nil,
			}
			l3 = l3.Next

			l1 = l1.Next
		} else if l2 != nil {
			value = l2.Val + carry
			carry = value / 10
			value = value % 10

			l3.Next = &ListNode{
				Val:  value,
				Next: nil,
			}
			l3 = l3.Next

			l2 = l2.Next
		}
	}

	if carry > 0 {

		l3.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return output
}

func print(node *ListNode) {
	for node != nil {
		fmt.Println("Val is ", node.Val)
		node = node.Next
	}
}
