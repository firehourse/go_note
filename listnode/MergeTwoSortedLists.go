package listnode

/**
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 首先他會給定兩個 sorted linked list 也就是說他已經幫忙排好序了
	// 那我們只需要宣告一個鏈表來對她進行裝箱其實就可以了
	// 先宣告一個dummy node
	dummy := &ListNode{}
	// 接著宣告一個head node
	head := dummy
	// 至少要走完一個鏈表
	for list1 != nil && list2 != nil {
		// 比大小
		if list1.Val < list2.Val {
			// 用head操作把地址塞進dummy
			head.Next = list1
			// 指向下一個位置
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	// 當其中一個走完另一個還沒走完就可以直接串起來
	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}
	// dummy還指著虛擬頭
	return dummy.Next
}
