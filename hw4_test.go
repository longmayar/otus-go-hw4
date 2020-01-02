package hw4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListPushFront(t *testing.T) {
	list := List{}
	list.PushFront(3)
	list.PushFront(2)
	list.PushFront(1)

	var actual []int
	expected := []int{1,2,3}

	item := list.First()
	for item != nil {
		actual = append(actual, item.Value())
		item = item.next
	}

	assert.Equal(t, expected, actual)
}

func TestListPushBack(t *testing.T) {
	list := List{}
	list.PushBack(3)
	list.PushBack(2)
	list.PushBack(1)

	var actual []int
	expected := []int{3,2,1}

	item := list.First()
	for item != nil {
		actual = append(actual, item.Value())
		item = item.next
	}

	assert.Equal(t, expected, actual)
}

func TestListLen(t *testing.T) {
	list := List{}
	list.PushFront(1)
	list.PushBack(1)
	list.PushFront(1)
	list.PushBack(1)
	list.PushFront(1)
	list.PushBack(1)

	assert.Equal(t, list.Len(), 6)
}

func TestListFirst(t *testing.T) {
	list := List{}
	list.PushFront(3)
	list.PushFront(2)
	list.PushFront(1)

	actual := list.First().value
	expected := 1

	assert.Equal(t, expected, actual)
}

func TestListLast(t *testing.T) {
	list := List{}
	list.PushFront(3)
	list.PushFront(2)
	list.PushFront(1)

	actual := list.Last().value
	expected := 3

	assert.Equal(t, expected, actual)
}

func TestListRemove(t *testing.T) {
	list := List{}
	list.PushBack(5)
	list.PushBack(4)
	list.PushBack(3)
	list.PushBack(2)
	list.PushBack(1)

	itemDelOrphan := &Item{value: 0}
	itemDel5 := list.First()
	itemDel1 := list.Last()

	itemDel3 := list.First()
	for itemDel3 != nil {
		if itemDel3.Value() == 3 {
			break
		}
		itemDel3 = itemDel3.next
	}

	list.Remove(itemDelOrphan)
	list.Remove(itemDel5)
	list.Remove(itemDel3)
	list.Remove(itemDel3)
	list.Remove(itemDel3)
	list.Remove(itemDel3)
	list.Remove(itemDel1)

	var actual []int
	expected := []int{4,2}

	item := list.First()
	for item != nil {
		actual = append(actual, item.Value())
		item = item.next
	}

	expectedLen := 2
	actualLen := list.Len()

	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedLen, actualLen)
}

func TestItemValue(t *testing.T) {
	item := &Item{value: 1}
	actual := item.Value()
	expected := 1

	assert.Equal(t, expected, actual)
}

func TestItemPrev(t *testing.T) {
	prevItem := &Item{value: 1}
	item := &Item{value: 1}
	item.prev = prevItem

	actual := item.Prev()
	expected := prevItem

	assert.Equal(t, expected, actual)
}

func TestItemNext(t *testing.T) {
	nextItem := &Item{value: 3}
	item := &Item{value: 1}
	item.next = nextItem

	actual := item.Next()
	expected := nextItem

	assert.Equal(t, expected, actual)
}