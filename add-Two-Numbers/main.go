/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := false
	var result *ListNode = nil //:= ListNode{0, nil}
	currentDigit := &ListNode{0, nil}
	for {
		var leftCurrentDigit int
		var rightCurrentDigit int
		if l1 != nil {
			leftCurrentDigit = l1.Val
		}
		if l2 != nil {
			rightCurrentDigit = l2.Val
		}
		digitResult := leftCurrentDigit + rightCurrentDigit
		if carry {
			carry = false
			digitResult++
		}
		if digitResult >= 10 {
			carry = true
			digitResult -= 10
		}
		currentDigit.Val = digitResult
		if result == nil {
			result = currentDigit
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil {
			if carry {
				currentDigit.Next = &ListNode{1, nil}
			}
			break
		}
		currentDigit.Next = &ListNode{0, nil}
		currentDigit = currentDigit.Next
	}
	return result
}
