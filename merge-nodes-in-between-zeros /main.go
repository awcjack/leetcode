/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}

	resultArray := []int{}
	counter := 0
	for head != nil {
		if head.Val == 0 {
			if counter != 0 {
				resultArray = append(resultArray, counter)
				counter = 0
			}
			head = head.Next
		} else {
			counter += head.Val
			head = head.Next
		}
	}

	if len(resultArray) > 0 {
		resultHead := &ListNode{resultArray[0], nil}
		head = resultHead
		for i := 1; i < len(resultArray); i++ {
			head.Next = &ListNode{resultArray[i], nil}
			head = head.Next
		}
		return resultHead
	}
	return nil
}
