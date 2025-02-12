package validParantheses

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	//list1 := &ListNode{}
	//list2 := &ListNode{Val: 1}
	//list3 := &ListNode{Val: 2}

	lists := []*ListNode{nil}
	//lists := []*ListNode{} // Input array
	//var expectedResult = &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}} // The expected answer with correct length.
	var expectedResult1 = []*ListNode{nil}
	res := mergeKLists(lists) // Calls your implementation

	assert.Equal(t, expectedResult1, res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	valuesSlice := []int{}
	for _, v := range lists {
		if v == nil {
			continue
		}
		for v.Next != nil {
			valuesSlice = append(valuesSlice, v.Val)
			v = v.Next
		}
		valuesSlice = append(valuesSlice, v.Val)
	}
	if len(valuesSlice) == 0 {
		return nil
	}
	slices.Sort(valuesSlice)
	return arrayToList(valuesSlice)
}
func arrayToList(array []int) *ListNode {
	head := &ListNode{Val: array[0], Next: nil}
	curr := head
	for i := 1; i < len(array); i++ {
		curr.Next = &ListNode{Val: array[i], Next: nil}
		curr = curr.Next
	}
	return head
}
