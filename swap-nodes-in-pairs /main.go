type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := head
	head = head.Next
	currentTemp := temp
	temp = temp.Next
	currentTemp.Next = temp.Next
	temp.Next = currentTemp
	for {
		if currentTemp.Next != nil && currentTemp.Next.Next != nil {
			temp = currentTemp.Next
			currentTemp.Next = currentTemp.Next.Next
			temp.Next = currentTemp.Next.Next
			currentTemp.Next.Next = temp
			currentTemp = currentTemp.Next.Next
		} else {
			break
		}
	}

	return head
}

func main() {
	head := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	newHead := swapPairs(&head)
	for {
		fmt.Println(newHead)
		if newHead.Next != nil {
			newHead = newHead.Next
		} else {
			break
		}
	}
}
