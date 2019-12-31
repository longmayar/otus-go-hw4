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

func TestListRemove(t *testing.T) {
	list := List{}
	list.PushBack(3)
	list.PushBack(2)
	list.PushBack(1)

	list.Remove(list.Last())

	var actual []int
	expected := []int{3,2}

	item := list.First()
	for item != nil {
		actual = append(actual, item.Value())
		item = item.next
	}

	assert.Equal(t, expected, actual)
}

func TestListFirst(t *testing.T) {
	list := List{}
	list.PushFront(3)
	list.PushFront(2)
	list.PushFront(1)

	actual := list.First().value
	expected := NewItem(1).value

	assert.Equal(t, expected, actual)
}

func TestListLast(t *testing.T) {
	list := List{}
	list.PushFront(3)
	list.PushFront(2)
	list.PushFront(1)

	actual := list.Last().value
	expected := NewItem(3).value

	assert.Equal(t, expected, actual)
}

func TestItemValue(t *testing.T) {
	item := NewItem(1)
	actual := item.Value()
	expected := 1

	assert.Equal(t, expected, actual)
}

func TestItemPrev(t *testing.T) {
	prevItem := NewItem(1)
	item := NewItem(1)
	item.prev = prevItem

	actual := item.Prev()
	expected := prevItem

	assert.Equal(t, expected, actual)
}

func TestItemNext(t *testing.T) {
	nextItem := NewItem(3)
	item := NewItem(1)
	item.next = nextItem

	actual := item.Next()
	expected := nextItem

	assert.Equal(t, expected, actual)
}