/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	hashMap := make(map[int]int)
	for {
		if list1 == nil && list2 == nil {
			break
		}
		if list1 != nil {
			hashMap[list1.Val]++
			list1 = list1.Next
		}
		if list2 != nil {
			hashMap[list2.Val]++
			list2 = list2.Next
		}
	}
	keys := make([]int, 0, len(hashMap))
	for k := range hashMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var result *ListNode
	var pointer *ListNode
	for _, k := range keys {
		for i := 0; i < hashMap[k]; i++ {
			if pointer == nil {
				pointer = &ListNode{k, nil}
				result = pointer
			} else {
				pointer.Next = &ListNode{k, nil}
				pointer = pointer.Next
			}
		}
	}
	return result
}
